package bootstrappers

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/duc-cnzj/mars-client/v4/types"
	"github.com/duc-cnzj/mars/frontend"
	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/contracts"
	e "github.com/duc-cnzj/mars/internal/event/events"
	"github.com/duc-cnzj/mars/internal/grpc/services"
	"github.com/duc-cnzj/mars/internal/middlewares"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/internal/models"
	"github.com/duc-cnzj/mars/internal/socket"
	"github.com/duc-cnzj/mars/internal/utils"
	"github.com/duc-cnzj/mars/third_party/doc/data"
	swagger_ui "github.com/duc-cnzj/mars/third_party/doc/swagger-ui"

	"github.com/dustin/go-humanize"
	"github.com/gorilla/mux"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
)

type ApiGatewayBootstrapper struct{}

func (a *ApiGatewayBootstrapper) Bootstrap(app contracts.ApplicationInterface) error {
	app.AddServer(&apiGateway{endpoint: fmt.Sprintf("localhost:%s", app.Config().GrpcPort)})
	app.RegisterAfterShutdownFunc(func(app contracts.ApplicationInterface) {
		t := time.NewTimer(5 * time.Second)
		defer t.Stop()
		ch := make(chan struct{}, 1)
		go func() {
			socket.Wait.Wait()
			ch <- struct{}{}
		}()
		select {
		case <-ch:
			mlog.Info("[Websocket]: all socket connection closed")
		case <-t.C:
			mlog.Warningf("[Websocket]: 等待超时, 未等待所有 socket 连接退出，当前剩余连接 %v 个。", socket.Wait.Count())
		}
	})

	return nil
}

type apiGateway struct {
	endpoint string
	server   *http.Server
}

func (a *apiGateway) Run(ctx context.Context) error {
	mlog.Infof("[Server]: start apiGateway runner at %s.", a.endpoint)

	router := mux.NewRouter()

	gmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:  true,
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}))

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithFilterFunc(middlewares.TracingIgnoreFn), grpc_opentracing.WithTracer(opentracing.GlobalTracer()))),
	}

	for _, f := range services.RegisteredEndpoints() {
		if err := f(ctx, gmux, a.endpoint, opts); err != nil {
			return err
		}
	}

	handFile(gmux)
	handleDownloadConfig(gmux)
	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.Write([]byte("pong"))
	})
	serveWs(router)
	frontend.LoadFrontendRoutes(router)
	LoadSwaggerUI(router)
	router.PathPrefix("/").Handler(gmux)

	s := &http.Server{
		Addr: ":" + app.Config().AppPort,
		Handler: middlewares.TracingWrapper(
			middlewares.RouteLogger(
				middlewares.AllowCORS(
					router,
				),
			),
		),
	}

	a.server = s

	go func(s *http.Server) {
		mlog.Info("api-gateway start at ", s.Addr)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mlog.Error(err)
		}
	}(s)

	return nil
}

func (a *apiGateway) Shutdown(ctx context.Context) error {
	defer mlog.Info("[Server]: shutdown api-gateway runner.")
	if a.server == nil {
		return nil
	}

	return a.server.Shutdown(ctx)
}

func handFile(gmux *runtime.ServeMux) {
	gmux.HandlePath("POST", "/api/files", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		if req, ok := authenticated(r); ok {
			handleBinaryFileUpload(w, req)
			return
		}
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
	})
	gmux.HandlePath("GET", "/api/download_file/{id}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		idstr, ok := pathParams["id"]
		if !ok {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idstr)
		if err != nil {
			http.Error(w, "bad id", http.StatusBadRequest)
			return
		}
		if req, ok := authenticated(r); ok {
			handleDownload(w, req, id)
			return
		}
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
	})
	gmux.HandlePath("GET", "/api/raw_file/{id}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		idstr, ok := pathParams["id"]
		if !ok {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}
		if _, ok := authenticated(r); ok {
			var f models.File
			if err := app.DB().First(&f, idstr).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					http.Error(w, "not found", http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			file, err := os.ReadFile(f.Path)
			if err == nil {
				w.Write(file)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
	})
}

func handleDownload(w http.ResponseWriter, r *http.Request, fid int) {
	var fil = &models.File{ID: fid}
	if err := app.DB().First(&fil).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	fileName := filepath.Base(fil.Path)

	user := r.Context().Value(authCtx{}).(*contracts.UserInfo)
	e.AuditLog(user.Name,
		types.EventActionType_Download,
		fmt.Sprintf("下载文件 '%s', 大小 %s",
			fil.Path, humanize.Bytes(fil.Size)), nil, nil)
	open, err := os.Open(fil.Path)
	if err != nil {
		mlog.Error(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	defer open.Close()

	download(w, fileName, open)
}

func download(w http.ResponseWriter, filename string, reader io.Reader) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, url.QueryEscape(filename)))
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Access-Control-Expose-Headers", "*")

	if _, err := io.Copy(w, bufio.NewReaderSize(reader, 1024*1024*5)); err != nil {
		mlog.Error(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func authenticated(r *http.Request) (*http.Request, bool) {
	if verifyToken, b := app.Auth().VerifyToken(r.Header.Get("Authorization")); b {
		return r.WithContext(context.WithValue(r.Context(), authCtx{}, &verifyToken.UserInfo)), true
	}

	return nil, false
}

func handleBinaryFileUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(int64(app.Config().MaxUploadSize())); err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	f, h, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer f.Close()

	info := r.Context().Value(authCtx{}).(*contracts.UserInfo)

	var uploader contracts.Uploader = app.Uploader()
	// 某个用户/那天/时间/文件名称
	put, err := uploader.Disk("users").Put(
		fmt.Sprintf("%s/%s/%s/%s",
			info.Name,
			time.Now().Format("2006-01-02"),
			fmt.Sprintf("%s-%s", time.Now().Format("15-04-05"), utils.RandomString(20)),
			h.Filename), f)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to upload file %s", err.Error()), http.StatusInternalServerError)
		return
	}

	file := models.File{Path: put.Path(), Username: info.Name, Size: put.Size()}
	app.DB().Create(&file)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var res = struct {
		ID int `json:"id"`
	}{
		ID: file.ID,
	}
	marshal, _ := json.Marshal(&res)
	w.Write(marshal)
}

func serveWs(mux *mux.Router) {
	ws := socket.NewWebsocketManager()
	ws.TickClusterHealth()
	mux.HandleFunc("/api/ws_info", ws.Info)
	mux.HandleFunc("/ws", ws.Ws)
}

type ExportProject struct {
	DefaultBranch string `json:"default_branch"`
	Name          string `json:"name"`
	GitProjectId  int    `json:"git_project_id"`
	Enabled       bool   `json:"enabled"`
	GlobalEnabled bool   `json:"global_enabled"`
	GlobalConfig  string `json:"global_config"`
}

func handleDownloadConfig(gmux *runtime.ServeMux) {
	gmux.HandlePath("GET", "/api/config/export", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		req, ok := authenticated(r)
		if !ok {
			http.Error(w, "Unauthenticated", http.StatusUnauthorized)
			return
		}
		user := req.Context().Value(authCtx{}).(*contracts.UserInfo)
		if !user.IsAdmin() {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		var projects []models.GitProject
		app.DB().Find(&projects)
		data := make([]ExportProject, 0, len(projects))
		for _, gitProject := range projects {
			data = append(data, ExportProject{
				DefaultBranch: gitProject.DefaultBranch,
				Name:          gitProject.Name,
				GitProjectId:  gitProject.GitProjectId,
				Enabled:       gitProject.Enabled,
				GlobalEnabled: gitProject.GlobalEnabled,
				GlobalConfig:  gitProject.GlobalConfig,
			})
		}
		marshal, _ := json.MarshalIndent(&data, "", "\t")
		e.AuditLog(user.Name,
			types.EventActionType_Download,
			"下载配置文件", nil, &e.StringYamlPrettier{Str: string(marshal)})
		download(w, "mars-config.json", bytes.NewReader(marshal))
	})
	gmux.HandlePath("POST", "/api/config/import", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		req, ok := authenticated(r)
		if !ok {
			http.Error(w, "Unauthenticated", http.StatusUnauthorized)
			return
		}
		user := req.Context().Value(authCtx{}).(*contracts.UserInfo)
		if !user.IsAdmin() {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		if err := r.ParseMultipartForm(int64(app.Config().MaxUploadSize())); err != nil {
			http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
			return
		}

		f, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
			return
		}
		defer f.Close()
		all, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		e.AuditLog(user.Name,
			types.EventActionType_Upload,
			"导入配置文件", nil, &e.StringYamlPrettier{Str: string(all)})
		var data []ExportProject
		err = json.Unmarshal(all, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, item := range data {
			var p models.GitProject
			if err := app.DB().Where("`git_project_id` = ?", item.GitProjectId).First(&p).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				app.DB().Create(&models.GitProject{
					DefaultBranch: item.DefaultBranch,
					Name:          item.Name,
					GitProjectId:  item.GitProjectId,
					Enabled:       item.Enabled,
					GlobalEnabled: item.GlobalEnabled,
					GlobalConfig:  item.GlobalConfig,
				})
			} else {
				app.DB().Model(&p).Updates(map[string]any{
					"default_branch": item.DefaultBranch,
					"name":           item.Name,
					"git_project_id": item.GitProjectId,
					"enabled":        item.Enabled,
					"global_enabled": item.GlobalEnabled,
					"global_config":  item.GlobalConfig,
				})
			}
		}
		w.WriteHeader(204)
		w.Write([]byte(""))
	})
}

func LoadSwaggerUI(mux *mux.Router) {
	subrouter := mux.PathPrefix("").Subrouter()
	subrouter.Use(middlewares.HttpCache)

	subrouter.Handle("/doc/swagger.json",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(data.SwaggerJson)
		}),
	)

	subrouter.PathPrefix("/docs/").Handler(
		http.StripPrefix("/docs/", http.FileServer(http.FS(swagger_ui.SwaggerUI))),
	)
}

type authCtx struct{}
