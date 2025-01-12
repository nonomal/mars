package services

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	clientgoexec "k8s.io/client-go/util/exec"

	"github.com/duc-cnzj/mars-client/v4/container"
	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/internal/models"
	"github.com/duc-cnzj/mars/internal/utils"
)

func init() {
	RegisterServer(func(s grpc.ServiceRegistrar, app contracts.ApplicationInterface) {
		container.RegisterContainerServer(s, new(Container))
	})
	RegisterEndpoint(container.RegisterContainerHandlerFromEndpoint)
}

type Container struct {
	container.UnsafeContainerServer
}

func (c *Container) IsPodRunning(_ context.Context, request *container.IsPodRunningRequest) (*container.IsPodRunningResponse, error) {
	running, reason := utils.IsPodRunning(request.GetNamespace(), request.GetPod())

	return &container.IsPodRunningResponse{Running: running, Reason: reason}, nil
}

func (c *Container) IsPodExists(_ context.Context, request *container.IsPodExistsRequest) (*container.IsPodExistsResponse, error) {
	_, err := app.K8sClientSet().CoreV1().Pods(request.Namespace).Get(context.TODO(), request.Pod, metav1.GetOptions{})
	if err != nil && apierrors.IsNotFound(err) {
		return &container.IsPodExistsResponse{Exists: false}, nil
	}

	return &container.IsPodExistsResponse{Exists: true}, nil
}

type exitCodeStatus struct {
	message string
	code    int
}

func (c *Container) Exec(request *container.ExecRequest, server container.Container_ExecServer) error {
	running, reason := utils.IsPodRunning(request.Namespace, request.Pod)
	if !running {
		return errors.New(reason)
	}

	if request.Container == "" {
		pod, _ := app.K8sClientSet().CoreV1().Pods(request.Namespace).Get(context.TODO(), request.Pod, metav1.GetOptions{})
		request.Container = FindDefaultContainer(pod)
		mlog.Debug("使用默认的容器: ", request.Container)
	}

	peo := &v1.PodExecOptions{
		Stdin:     false,
		Stdout:    true,
		Stderr:    true,
		TTY:       false,
		Container: request.Container,
		Command:   request.Command,
	}

	req := app.K8sClient().Client.CoreV1().
		RESTClient().
		Post().
		Namespace(request.Namespace).
		Resource("pods").
		SubResource("exec").
		Name(request.Pod)

	params := req.VersionedParams(peo, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(app.K8sClient().RestConfig, "POST", params.URL())
	if err != nil {
		return err
	}
	var exitCode atomic.Value
	writer := newExecWriter()
	defer writer.Close()

	go func() {
		defer writer.Close()
		defer utils.HandlePanic("Exec")
		err := exec.Stream(remotecommand.StreamOptions{
			Stdin:             nil,
			Stdout:            writer,
			Stderr:            writer,
			Tty:               false,
			TerminalSizeQueue: nil,
		})
		if err != nil {
			if exitError, ok := err.(clientgoexec.ExitError); ok && exitError.Exited() {
				mlog.Debugf("[Container]: exit %v, exit code: %d, err: %v", exitError.Exited(), exitError.ExitStatus(), exitError.Error())
				exitCode.Store(&exitCodeStatus{
					message: exitError.Error(),
					code:    exitError.ExitStatus(),
				})
			} else {
				mlog.Error(err)
				exitCode.Store(&exitCodeStatus{
					message: err.Error(),
					code:    1,
				})
			}
		}
	}()

	for {
		select {
		case msg, ok := <-writer.ch:
			if !ok {
				ec := exitCode.Load()
				if ec != nil {
					ecs := ec.(*exitCodeStatus)
					server.Send(&container.ExecResponse{
						Error: &container.ExecError{
							Code:    int64(ecs.code),
							Message: ecs.message,
						},
					})
				}
				return nil
			}
			if err := server.Send(&container.ExecResponse{
				Message: msg,
			}); err != nil {
				return err
			}
		case <-server.Context().Done():
			return server.Context().Err()
		}
	}
}

func (c *Container) CopyToPod(ctx context.Context, request *container.CopyToPodRequest) (*container.CopyToPodResponse, error) {
	if running, reason := utils.IsPodRunning(request.Namespace, request.Pod); !running {
		return nil, status.Error(codes.NotFound, reason)
	}

	var file models.File
	if err := app.DB().First(&file, request.FileId).Error; err != nil {
		return nil, err
	}
	res, err := utils.CopyFileToPod(request.Namespace, request.Pod, request.Container, file.Path, "")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	app.DB().Model(&file).Updates(map[string]any{
		"namespace":      request.Namespace,
		"pod":            request.Pod,
		"container":      request.Container,
		"container_path": res.ContainerPath,
	})

	FileAuditLog(MustGetUser(ctx).Name,
		fmt.Sprintf("上传文件到 pod: %s/%s/%s, 容器路径: '%s', 大小: %s。",
			request.Namespace,
			request.Pod,
			request.Container,
			res.ContainerPath,
			humanize.Bytes(file.Size),
		), file.ID)

	return &container.CopyToPodResponse{
		PodFilePath: res.TargetDir,
		Output:      res.ErrOut,
		FileName:    res.FileName,
	}, err
}

func (c *Container) StreamCopyToPod(server container.Container_StreamCopyToPodServer) error {
	var (
		fpath         string
		namespace     string
		pod           string
		containerName string
		user          = MustGetUser(server.Context())
		f             *os.File
		disk          = "grpc_upload"
	)
	defer f.Close()

	for {
		recv, err := server.Recv()
		if err != nil {
			if err == io.EOF && f != nil {
				stat, _ := f.Stat()
				f.Close()

				file := models.File{Path: f.Name(), Username: user.Name, Size: uint64(stat.Size())}
				app.DB().Create(&file)
				res, err := c.CopyToPod(server.Context(), &container.CopyToPodRequest{
					FileId:    int64(file.ID),
					Namespace: namespace,
					Pod:       pod,
					Container: containerName,
				})
				if err != nil {
					return err
				}
				return server.SendAndClose(&container.StreamCopyToPodResponse{
					Size:        stat.Size(),
					PodFilePath: res.PodFilePath,
					Output:      res.Output,
					Pod:         pod,
					Namespace:   namespace,
					Container:   containerName,
					Filename:    res.FileName,
				})
			}
			if f != nil {
				f.Close()
				app.Uploader().Disk(disk).Delete(f.Name())
			}
			return err
		}
		if fpath == "" {
			pod = recv.Pod
			namespace = recv.Namespace
			if recv.Container == "" {
				pod, err := app.K8sClientSet().CoreV1().Pods(recv.Namespace).Get(context.TODO(), recv.Pod, metav1.GetOptions{})
				if err != nil {
					return err
				}

				recv.Container = FindDefaultContainer(pod)
				mlog.Debug("使用默认的容器: ", recv.Container)
			}
			containerName = recv.Container
			running, reason := utils.IsPodRunning(recv.Namespace, recv.Pod)
			if !running {
				return errors.New(reason)
			}

			// 某个用户/那天/时间/文件名称
			// duc/2006-01-02/15-03-04-random-str/xxx.tgz
			p := fmt.Sprintf("%s/%s/%s/%s",
				user.Name,
				time.Now().Format("2006-01-02"),
				fmt.Sprintf("%s-%s", time.Now().Format("15-04-05"), utils.RandomString(20)),
				filepath.Base(recv.GetFileName()))
			fpath = app.Uploader().Disk(disk).AbsolutePath(p)
			err := app.Uploader().Disk(disk).MkDir(filepath.Dir(p), true)
			if err != nil {
				mlog.Error(err)
				return err
			}
			f, err = os.Create(fpath)
			if err != nil {
				mlog.Error(err)
				return err
			}
		}

		f.Write(recv.GetData())
	}
}

func (c *Container) ContainerLog(ctx context.Context, request *container.LogRequest) (*container.LogResponse, error) {
	if running, reason := utils.IsPodRunning(request.Namespace, request.Pod); !running {
		return nil, status.Errorf(codes.NotFound, reason)
	}

	var limit int64 = 2000
	logs := app.K8sClientSet().CoreV1().Pods(request.Namespace).GetLogs(request.Pod, &v1.PodLogOptions{
		Container: request.Container,
		TailLines: &limit,
	})
	do := logs.Do(context.Background())
	raw, err := do.Raw()
	if err != nil {
		return nil, err
	}

	return &container.LogResponse{
		Namespace:     request.Namespace,
		PodName:       request.Pod,
		ContainerName: request.Container,
		Log:           string(raw),
	}, nil
}

func (c *Container) StreamContainerLog(request *container.LogRequest, server container.Container_StreamContainerLogServer) error {
	if running, reason := utils.IsPodRunning(request.Namespace, request.Pod); !running {
		return status.Errorf(codes.NotFound, reason)
	}

	var limit int64 = 2000
	logs := app.K8sClientSet().CoreV1().Pods(request.Namespace).GetLogs(request.Pod, &v1.PodLogOptions{
		Follow:    true,
		Container: request.Container,
		TailLines: &limit,
	})
	stream, err := logs.Stream(context.TODO())
	if err != nil {
		return err
	}
	bf := bufio.NewReader(stream)

	ch := make(chan []byte)
	go func() {
		defer func() {
			mlog.Debug("[Stream]:  read exit!")
			close(ch)
		}()
		defer utils.HandlePanic("StreamContainerLog")

		for {
			bytes, err := bf.ReadBytes('\n')
			if err != nil {
				mlog.Debugf("[Stream]: %v", err)
				return
			}
			ch <- bytes
		}
	}()

	for {
		select {
		case <-app.App().Done():
			stream.Close()
			err := errors.New("server shutdown")
			mlog.Debug("[Stream]: client exit with: ", err)
			return err
		case <-server.Context().Done():
			stream.Close()
			mlog.Debug("[Stream]: client exit with: ", server.Context().Err())
			return server.Context().Err()
		case msg, ok := <-ch:
			if !ok {
				stream.Close()
				return errors.New("[Stream]: channel close")
			}

			if err := server.Send(&container.LogResponse{
				Namespace:     request.Namespace,
				PodName:       request.Pod,
				ContainerName: request.Container,
				Log:           string(msg),
			}); err != nil {
				stream.Close()
				return err
			}
		}
	}
}

type closeable struct {
	sync.RWMutex
	closed bool
}

func (c *closeable) IsClosed() bool {
	c.RLock()
	defer c.RUnlock()
	return c.closed
}

func (c *closeable) Close() {
	c.Lock()
	defer c.Unlock()
	c.closed = true
}

type execWriter struct {
	reader io.Reader
	state  *closeable
	ch     chan string
}

func (rw *execWriter) IsClosed() bool {
	return rw.state.IsClosed()
}

func (rw *execWriter) Close() error {
	if rw.IsClosed() {
		return nil
	}
	rw.state.Close()
	close(rw.ch)

	return nil
}

func newExecWriter() *execWriter {
	return &execWriter{
		state: &closeable{},
		ch:    make(chan string, 100),
	}
}

func (rw *execWriter) Write(p []byte) (int, error) {
	if rw.state.IsClosed() {
		mlog.Warning("execWriter close")
		return 0, errors.New("closed")
	}
	rw.ch <- string(p)
	return len(p), nil
}

const DefaultContainerAnnotationName = "kubectl.kubernetes.io/default-container"

func FindDefaultContainer(pod *v1.Pod) string {
	if name := pod.Annotations[DefaultContainerAnnotationName]; len(name) > 0 {
		return name
	}

	for _, co := range pod.Spec.Containers {
		return co.Name
	}

	return ""
}
