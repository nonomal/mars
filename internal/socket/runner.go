package socket

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"

	"github.com/duc-cnzj/mars/api/v5/mars"
	"github.com/duc-cnzj/mars/api/v5/types"
	websocket_pb "github.com/duc-cnzj/mars/api/v5/websocket"
	"github.com/duc-cnzj/mars/v5/internal/application"
	"github.com/duc-cnzj/mars/v5/internal/auth"
	"github.com/duc-cnzj/mars/v5/internal/data"
	"github.com/duc-cnzj/mars/v5/internal/locker"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/repo"
	"github.com/duc-cnzj/mars/v5/internal/transformer"
	"github.com/duc-cnzj/mars/v5/internal/uploader"
	"github.com/duc-cnzj/mars/v5/internal/util/pipeline"
	"github.com/duc-cnzj/mars/v5/internal/util/rand"
	"github.com/duc-cnzj/mars/v5/internal/util/timer"
	mysort "github.com/duc-cnzj/mars/v5/internal/util/xsort"
	yaml2 "github.com/duc-cnzj/mars/v5/internal/util/yaml"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/release"
)

var ErrorVersionNotMatched = errors.New("[部署冲突]: 1. 可能是多个人同时部署导致 2. 项目已经存在")

const (
	ResultError             = websocket_pb.ResultType_Error
	ResultSuccess           = websocket_pb.ResultType_Success
	ResultDeployed          = websocket_pb.ResultType_Deployed
	ResultDeployFailed      = websocket_pb.ResultType_DeployedFailed
	ResultDeployCanceled    = websocket_pb.ResultType_DeployedCanceled
	ResultLogWithContainers = websocket_pb.ResultType_LogWithContainers

	WsSetUid             = websocket_pb.Type_SetUid
	WsReloadProjects     = websocket_pb.Type_ReloadProjects
	WsCancel             = websocket_pb.Type_CancelProject
	WsCreateProject      = websocket_pb.Type_CreateProject
	WsUpdateProject      = websocket_pb.Type_UpdateProject
	WsProcessPercent     = websocket_pb.Type_ProcessPercent
	WsClusterInfoSync    = websocket_pb.Type_ClusterInfoSync
	WsInternalError      = websocket_pb.Type_InternalError
	WsHandleExecShell    = websocket_pb.Type_HandleExecShell
	WsHandleExecShellMsg = websocket_pb.Type_HandleExecShellMsg
	WsHandleCloseShell   = websocket_pb.Type_HandleCloseShell
	WsAuthorize          = websocket_pb.Type_HandleAuthorize
	ProjectPodEvent      = websocket_pb.Type_ProjectPodEvent

	// Maximum message size allowed from peer.
	maxMessageSize = 1024 * 1024 * 20 // 20MB
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 8) / 10
)

var _ JobManager = (*jobManager)(nil)

type JobManager interface {
	//	创建一个新的Job
	NewJob(input *JobInput) Job
}

type Job interface {
	Stop(error)
	IsNotDryRun() bool

	ID() string
	GlobalLock() Job
	Validate() Job
	LoadConfigs() Job
	Run(ctx context.Context) Job
	Finish() Job
	Error() error
	Project() *repo.Project
	Manifests() []string

	OnError(p int, fn func(err error, sendResultToUser func())) Job
	OnSuccess(p int, fn func(err error, sendResultToUser func())) Job
	OnFinally(p int, fn func(err error, sendResultToUser func())) Job
}

type jobManager struct {
	logger mlog.Logger
	data   data.Data

	timer            timer.Timer
	releaseInstaller ReleaseInstaller
	nsRepo           repo.NamespaceRepo
	projRepo         repo.ProjectRepo
	eventRepo        repo.EventRepo
	k8sRepo          repo.K8sRepo
	helmRepo         repo.HelmerRepo
	repoRepo         repo.RepoRepo

	locker       locker.Locker
	uploader     uploader.Uploader
	pluginManger application.PluginManger
}

func NewJobManager(
	data data.Data,
	timer timer.Timer,
	logger mlog.Logger,
	releaseInstaller ReleaseInstaller,
	repoRepo repo.RepoRepo,
	nsRepo repo.NamespaceRepo,
	projRepo repo.ProjectRepo,
	helmer repo.HelmerRepo,
	uploader uploader.Uploader,
	locker locker.Locker,
	k8sRepo repo.K8sRepo,
	eventRepo repo.EventRepo,
	pl application.PluginManger,
) JobManager {
	return &jobManager{
		timer:            timer,
		releaseInstaller: releaseInstaller,
		uploader:         uploader,
		repoRepo:         repoRepo,
		data:             data,
		logger:           logger,
		nsRepo:           nsRepo,
		projRepo:         projRepo,
		k8sRepo:          k8sRepo,
		pluginManger:     pl,
		helmRepo:         helmer,
		locker:           locker,
		eventRepo:        eventRepo,
	}
}

func (j *jobManager) NewJob(input *JobInput) Job {
	var timeoutSeconds int64 = int64(input.TimeoutSeconds)
	if timeoutSeconds == 0 {
		timeoutSeconds = int64(j.data.Config().InstallTimeout.Seconds())
	}
	jb := &jobRunner{
		installer:       j.releaseInstaller,
		logger:          j.logger.WithModule("socket/job"),
		nsRepo:          j.nsRepo,
		projRepo:        j.projRepo,
		repoRepo:        j.repoRepo,
		pluginMgr:       j.pluginManger,
		helmer:          j.helmRepo,
		locker:          j.locker,
		k8sRepo:         j.k8sRepo,
		eventRepo:       j.eventRepo,
		timer:           j.timer,
		uploader:        j.uploader,
		loaders:         defaultLoaders(),
		dryRun:          input.DryRun,
		input:           input,
		finallyCallback: mysort.PrioritySort[func(err error, next func())]{},
		errorCallback:   mysort.PrioritySort[func(err error, next func())]{},
		successCallback: mysort.PrioritySort[func(err error, next func())]{},
		deployResult:    &deployResult{},
		valuesOptions:   &values.Options{},
		messageCh:       NewSafeWriteMessageCh(j.logger, 100),
		messager:        input.Messager,
		user:            input.User,
		timeoutSeconds:  timeoutSeconds,
	}
	jb.stopCtx, jb.stopFn = context.WithCancelCause(context.TODO())

	return jb
}

type JobInput struct {
	Type        websocket_pb.Type
	NamespaceId int32
	Name        string
	RepoID      int32
	GitBranch   string
	GitCommit   string
	Config      string
	Atomic      *bool
	ExtraValues []*websocket_pb.ExtraValue
	Version     *int32
	ProjectID   int32

	TimeoutSeconds int32
	User           *auth.UserInfo
	DryRun         bool

	PubSub   application.PubSub `json:"-"`
	Messager DeployMsger        `json:"-"`
}

func (job *JobInput) Slug() string {
	return GetSlugName(job.NamespaceId, job.Name)
}

func (job *JobInput) IsNotDryRun() bool {
	return !job.DryRun
}

type jobRunner struct {
	// 这些属性在 new runner 的时候就已经初始化了
	logger          mlog.Logger
	nsRepo          repo.NamespaceRepo
	projRepo        repo.ProjectRepo
	repoRepo        repo.RepoRepo
	helmer          repo.HelmerRepo
	locker          locker.Locker
	k8sRepo         repo.K8sRepo
	eventRepo       repo.EventRepo
	messager        DeployMsger
	timeoutSeconds  int64
	uploader        uploader.Uploader
	pluginMgr       application.PluginManger
	installer       ReleaseInstaller
	messageCh       SafeWriteMessageChan
	deployResult    *deployResult
	loaders         []Loader
	input           *JobInput
	finallyCallback mysort.PrioritySort[func(err error, next func())]
	errorCallback   mysort.PrioritySort[func(err error, next func())]
	successCallback mysort.PrioritySort[func(err error, next func())]
	stopCtx         context.Context
	stopFn          func(error)
	dryRun          bool
	user            *auth.UserInfo
	timer           timer.Timer

	// 这些属性在执行的时候才会初始化
	// Validate 阶段被初始化
	isNew            bool
	ns               *repo.Namespace
	repo             *repo.Repo
	config           *mars.Config
	project          *repo.Project
	imagePullSecrets []string
	commit           application.Commit
	oldConf          repo.YamlPrettier

	// LoadConfigs 阶段被初始化

	// 1. ChartFileLoader 时加载
	chart *chart.Chart

	// 2. UserConfigLoader 时加载
	userConfigYaml string

	// 3. ElementsLoader(自定义配置) 时加载
	elementValues    []string
	finalExtraValues []*websocket_pb.ExtraValue

	// 4. SystemVariableLoader 时加载
	// chart 的 替换完所有 <.Var> 之后的 values.yaml 内容
	systemValuesYaml string
	// systemValuesYaml 注入的变量
	vars vars

	// 5. MergeValuesLoader 时加载
	// 把 values.yaml + elementValues + 自定义配置 合并后的结果
	valuesOptions *values.Options

	err error

	// 部署成功后的 manifest
	manifests []string
}

func (j *jobRunner) ID() string {
	return j.input.Slug()
}

func (j *jobRunner) IsNotDryRun() bool {
	return !j.dryRun
}

func (j *jobRunner) GlobalLock() Job {
	if j.HasError() {
		return j
	}
	releaseFn, acquired := j.locker.RenewalAcquire(j.ID(), 30, 20)
	if !acquired {
		return j.SetError(errors.New("正在部署中，请稍后再试"))
	}

	return j.OnFinally(-1, func(err error, sendResultToUser func()) {
		sendResultToUser()
		releaseFn()
	})
}

func (j *jobRunner) Validate() Job {
	var err error
	if !j.typeValidated() {
		return j.SetError(errors.New("type error: " + j.input.Type.String()))
	}

	j.messager.SendMsg("[start]: 收到请求，开始创建项目")
	j.messager.To(5)

	j.messager.SendMsg("[Check]: 校验名称空间...")

	j.ns, err = j.nsRepo.Show(context.TODO(), int(j.input.NamespaceId))
	if err != nil {
		return j.SetError(fmt.Errorf("[FAILED]: 校验名称空间: %w", err))
	}

	j.messager.SendMsg("[Loading]: 加载用户配置")
	j.messager.To(10)

	j.repo, err = j.repoRepo.Get(context.TODO(), int(j.input.RepoID))
	if err != nil {
		return j.SetError(err)
	}
	j.config = j.repo.MarsConfig

	j.messager.SendMsg("[Check]: 检查项目是否存在")

	found, err := j.projRepo.FindByName(context.TODO(), j.input.Name, j.ns.ID)
	if err != nil {
		createProjectInput := &repo.CreateProjectInput{
			Name:         j.input.Name,
			GitProjectID: int(j.repo.GitProjectID),
			GitBranch:    j.input.GitBranch,
			GitCommit:    j.input.GitCommit,
			Config:       j.input.Config,
			Atomic:       j.input.Atomic,
			ConfigType:   j.config.ConfigFileType,
			NamespaceID:  j.ns.ID,
			RepoID:       j.repo.ID,
			Creator:      j.user.Email,
		}
		j.messager.SendMsg("[Check]: 新建项目")
		createProjectInput.DeployStatus = types.Deploy_StatusDeploying
		j.isNew = true
		if j.IsNotDryRun() {
			j.project, err = j.projRepo.Create(context.TODO(), createProjectInput)
			if err != nil {
				j.logger.Warning(err)
				return j.SetError(err)
			}
			createdID := j.project.ID
			j.OnError(1, func(err error, sendResultToUser func()) {
				j.logger.Debug("清理项目")
				j.projRepo.Delete(context.TODO(), createdID)
				sendResultToUser()
			})
		}
	} else {
		j.project = found
		version := j.project.Version
		if j.IsNotDryRun() {
			j.messager.SendMsg(fmt.Sprintf("[Check]: 检查当前版本, version: %v", lo.FromPtr(j.input.Version)))
			j.project, err = j.projRepo.UpdateStatusByVersion(context.TODO(), int(j.input.ProjectID), types.Deploy_StatusDeploying, int(lo.FromPtr(j.input.Version)))
			if err != nil {
				return j.SetError(fmt.Errorf("%w: %w", ErrorVersionNotMatched, err))
			}
			j.OnError(1, func(err error, sendResultToUser func()) {
				j.project, _ = j.projRepo.UpdateVersion(context.TODO(), j.project.ID, version)
				sendResultToUser()
			})
		}
	}

	if j.IsNotDryRun() {
		reloadMessage := &websocket_pb.WsReloadProjectsResponse{
			Metadata:    &websocket_pb.Metadata{Type: WsReloadProjects},
			NamespaceId: int32(j.ns.ID),
		}
		j.PubSub().ToAll(reloadMessage)
		j.OnFinally(1, func(err error, sendResultToUser func()) {
			// 如果状态出现问题，只有拿到锁的才能更新状态
			j.project, _ = j.projRepo.UpdateDeployStatus(context.TODO(), j.project.ID, j.helmer.ReleaseStatus(j.Project().Name, j.ns.Name))
			j.PubSub().ToAll(reloadMessage)
			sendResultToUser()
		})
	}

	j.imagePullSecrets = j.ns.ImagePullSecrets
	j.commit = NewEmptyCommit()
	if j.repo.NeedGitRepo {
		j.commit, err = j.pluginMgr.Git().GetCommit(fmt.Sprintf("%d", j.repo.GitProjectID), j.input.GitCommit)
	}
	if !j.isNew {
		j.oldConf = toProjectEventYaml(j.project)
	}

	return j.SetError(err)
}

func (j *jobRunner) typeValidated() bool {
	return j.input.Type == websocket_pb.Type_CreateProject ||
		j.input.Type == websocket_pb.Type_UpdateProject ||
		j.input.Type == websocket_pb.Type_ApplyProject
}

func (j *jobRunner) LoadConfigs() Job {
	if j.HasError() {
		return j
	}
	eg, _ := errgroup.WithContext(j.stopCtx)
	eg.Go(func() error {
		defer j.logger.HandlePanic("LoadConfigs")
		return func() error {
			j.messager.SendMsg("[Check]: 加载项目文件")

			for _, defaultLoader := range j.loaders {
				if err := j.GetStoppedErrorIfHas(); err != nil {
					return err
				}
				if err := defaultLoader.Load(j); err != nil {
					return err
				}
			}

			return nil
		}()
	})

	return j.SetError(eg.Wait())
}

func (j *jobRunner) Run(ctx context.Context) Job {
	if j.HasError() {
		return j
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		defer j.logger.HandlePanic("[Websocket]: jobRunner Run")
		defer j.messageCh.Close()
		j.HandleMessage(ctx)
		return nil
	})

	eg.Go(func() error {
		defer j.logger.HandlePanic("[Websocket]: jobRunner Run")
		var (
			result *release.Release
			err    error
		)

		j.messager.SendMsg("worker 已就绪, 准备安装")
		if result, err = j.installer.Run(ctx, &InstallInput{
			IsNew:        j.isNew,
			Wait:         lo.FromPtr(j.input.Atomic),
			Chart:        j.chart,
			ValueOptions: j.valuesOptions,
			DryRun:       j.dryRun,
			ReleaseName:  j.project.Name,
			Namespace:    j.ns.Name,
			Description:  j.commit.GetTitle(),
			messageChan:  j.messageCh,
			percenter:    j.messager,
		}); err != nil {
			j.logger.Errorf("[Websocket]: %v", err)
			j.messageCh.Send(MessageItem{
				Msg:  err.Error(),
				Type: MessageError,
			})
			return err
		}

		coalesceValues, _ := chartutil.CoalesceValues(j.chart, result.Config)
		marshal, _ := yaml2.PrettyMarshal(&coalesceValues)
		manifests := j.k8sRepo.SplitManifests(result.Manifest)
		j.manifests = manifests
		var updateProjectInput = &repo.UpdateProjectInput{
			ID:           j.project.ID,
			GitBranch:    j.input.GitBranch,
			GitCommit:    j.input.GitCommit,
			Config:       j.input.Config,
			Atomic:       j.input.Atomic,
			ConfigType:   j.config.GetConfigFileType(),
			PodSelectors: j.k8sRepo.GetPodSelectorsByManifest(manifests),
			DockerImage: matchDockerImage(pipelineVars{
				Pipeline: j.vars.MustGetString("Pipeline"),
				Commit:   j.vars.MustGetString("Commit"),
				Branch:   j.vars.MustGetString("Branch"),
			}, result.Manifest),
			GitCommitTitle:   j.commit.GetTitle(),
			GitCommitWebURL:  j.commit.GetWebURL(),
			GitCommitAuthor:  j.commit.GetAuthorName(),
			GitCommitDate:    j.commit.GetCommittedDate(),
			ExtraValues:      j.input.ExtraValues,
			FinalExtraValues: j.finalExtraValues,
			EnvValues:        j.vars.ToKeyValue(),
			OverrideValues:   string(marshal),
			Manifest:         j.manifests,
		}

		var (
			oldConf repo.YamlPrettier = j.oldConf
			newConf repo.YamlPrettier
		)

		if j.IsNotDryRun() {
			j.project, err = j.projRepo.UpdateProject(context.TODO(), updateProjectInput)
			if err != nil {
				j.logger.Warning(err)
				return err
			}

			newConf = toProjectEventYaml(j.project)
			j.eventRepo.Dispatch(repo.EventProjectChanged, &repo.ProjectChangedData{
				ID:       j.project.ID,
				Username: j.user.Name,
			})
		}

		var act types.EventActionType = types.EventActionType_Create
		if !j.isNew {
			act = types.EventActionType_Update
		}
		if j.dryRun {
			act = types.EventActionType_DryRun
			prettyMarshal, _ := yaml2.PrettyMarshal(j.input)
			newConf = &repo.StringYamlPrettier{Str: string(prettyMarshal)}
		}
		j.eventRepo.AuditLogWithChange(
			act,
			j.user.Name,
			fmt.Sprintf("%s 项目: %s/%s", act.String(), j.ns.Name, j.Project().Name),
			oldConf, newConf)
		j.messager.To(100)
		j.messageCh.Send(MessageItem{
			Msg:  "部署成功",
			Type: MessageSuccess,
		})
		return nil
	})

	return j.SetError(eg.Wait())
}

func (j *jobRunner) Finish() Job {
	j.logger.Debug("finished")

	var callbacks []func(err error, next func())

	// Run error hooks
	if j.HasError() {
		func(err error) {
			pmodel := transformer.FromProject(j.project)
			j.deployResult.Set(websocket_pb.ResultType_DeployedFailed, err.Error(), pmodel)

			if e := j.GetStoppedErrorIfHas(); e != nil {
				j.deployResult.Set(websocket_pb.ResultType_DeployedCanceled, e.Error(), pmodel)
				err = e
			}
		}(j.Error())
		callbacks = append(callbacks, j.errorCallback.Sort()...)
	}

	// Run success hooks
	if !j.HasError() {
		callbacks = append(callbacks, j.successCallback.Sort()...)
	}

	// run finally hooks
	callbacks = append(callbacks, j.finallyCallback.Sort()...)

	pipeline.New[error]().
		Send(j.Error()).
		Through(callbacks...).
		Then(func(error) {
			if j.deployResult.IsSet() {
				j.messager.SendDeployedResult(j.deployResult.ResultType(), j.deployResult.Msg(), j.deployResult.Model())
			}
			j.logger.Debug("SendDeployedResult")
		})

	return j
}

func (j *jobRunner) Manifests() []string {
	return j.manifests
}

func (j *jobRunner) Stop(err error) {
	j.messager.SendMsg("收到取消信号, 开始停止部署~")
	j.logger.Debugf("stop deploy jobRunner, because '%v'", err)
	j.stopFn(err)
}

func (j *jobRunner) OnError(p int, fn func(err error, sendResultToUser func())) Job {
	j.errorCallback.Add(p, fn)
	return j
}

func (j *jobRunner) OnSuccess(p int, fn func(err error, sendResultToUser func())) Job {
	j.successCallback.Add(p, fn)
	return j
}

func (j *jobRunner) OnFinally(p int, fn func(err error, sendResultToUser func())) Job {
	j.finallyCallback.Add(p, fn)
	return j
}

func (j *jobRunner) Error() error {
	return j.err
}

func (j *jobRunner) SetError(err error) *jobRunner {
	j.err = err
	return j
}

func (j *jobRunner) HasError() bool {
	return j.err != nil
}

func (j *jobRunner) Project() *repo.Project {
	return j.project
}

func (j *jobRunner) PubSub() application.PubSub {
	return j.input.PubSub
}

func (j *jobRunner) IsStopped() bool {
	select {
	case <-j.stopCtx.Done():
		return true
	default:
	}

	return false
}

func (j *jobRunner) GetStoppedErrorIfHas() error {
	if j.IsStopped() {
		return context.Cause(j.stopCtx)
	}
	return nil
}

func (j *jobRunner) WriteConfigYamlToTmpFile(data []byte) (string, io.Closer, error) {
	file := fmt.Sprintf("mars-%s-%s.yaml", j.timer.Now().Format("2006-01-02"), rand.String(20))
	info, err := j.uploader.LocalUploader().Put(file, bytes.NewReader(data))
	if err != nil {
		return "", nil, err
	}
	path := info.Path()

	return path, NewCloser(func() error {
		j.logger.Debug("delete file: " + path)
		if err := j.uploader.LocalUploader().Delete(path); err != nil {
			j.logger.Error("WriteConfigYamlToTmpFile error: ", err)
			return err
		}

		return nil
	}), nil
}

func (j *jobRunner) DownloadFiles(pid any, commit string, files []string) (string, func(), error) {
	id := fmt.Sprintf("%v", pid)
	dir := fmt.Sprintf("mars_tmp_%s", rand.String(10))
	if err := j.uploader.LocalUploader().MkDir(dir, true); err != nil {
		return "", nil, err
	}

	return j.DownloadFilesToDir(id, commit, files, j.uploader.LocalUploader().AbsolutePath(dir))
}

func (j *jobRunner) DownloadFilesToDir(pid any, commit string, files []string, dir string) (string, func(), error) {
	wg := &sync.WaitGroup{}
	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			defer wg.Done()
			defer j.logger.HandlePanic("DownloadFilesToDir")
			raw, err := j.pluginMgr.Git().GetFileContentWithSha(fmt.Sprintf("%v", pid), commit, file)
			if err != nil {
				j.logger.Error(err)
			}
			localPath := filepath.Join(dir, file)
			if _, err := j.uploader.LocalUploader().Put(localPath, strings.NewReader(raw)); err != nil {
				j.logger.Errorf("[DownloadFilesToDir]: err '%s'", err.Error())
			}
		}(file)
	}
	wg.Wait()

	return dir, func() {
		err := j.uploader.LocalUploader().DeleteDir(dir)
		if err != nil {
			j.logger.Warning(err)
			return
		}
		j.logger.Debug("remove " + dir)
	}, nil
}

func (j *jobRunner) HandleMessage(ctx context.Context) {
	defer j.logger.Debug("HandleMessage exit")
	ch := j.messageCh.Chan()
	for {
		select {
		case <-ctx.Done():
			return
		case s, ok := <-ch:
			if !ok {
				return
			}
			switch s.Type {
			case MessageText:
				j.messager.SendMsgWithContainerLog(s.Msg, s.Containers)
			case MessageError:
				select {
				case <-j.stopCtx.Done():
					j.deployResult.Set(ResultDeployCanceled, context.Cause(j.stopCtx).Error(), transformer.FromProject(j.project))
				default:
					j.deployResult.Set(ResultDeployFailed, s.Msg, transformer.FromProject(j.project))
				}
				return
			case MessageSuccess:
				j.deployResult.Set(ResultDeployed, s.Msg, transformer.FromProject(j.project))
				return
			}
		}
	}
}

func toProjectEventYaml(p *repo.Project) repo.YamlPrettier {
	if p == nil {
		return nil
	}

	sort.Slice(p.EnvValues, func(i, j int) bool {
		return p.EnvValues[i].Key < p.EnvValues[j].Key
	})
	sort.Slice(p.ExtraValues, func(i, j int) bool {
		return p.ExtraValues[i].Path < p.ExtraValues[j].Path
	})
	sort.Slice(p.FinalExtraValues, func(i, j int) bool {
		return p.FinalExtraValues[i].Path < p.FinalExtraValues[j].Path
	})

	return repo.AnyYamlPrettier{
		"title":              p.GitCommitTitle,
		"branch":             p.GitBranch,
		"commit":             p.GitCommit,
		"atomic":             p.Atomic,
		"web_url":            p.GitCommitWebURL,
		"config":             p.Config,
		"env_values":         p.EnvValues,
		"extra_values":       p.ExtraValues,
		"final_extra_values": p.FinalExtraValues,
	}
}

type deployResult struct {
	sync.RWMutex
	result websocket_pb.ResultType
	msg    string
	model  *types.ProjectModel
	set    bool
}

func (d *deployResult) IsSet() bool {
	d.RLock()
	defer d.RUnlock()
	return d.set
}

func (d *deployResult) Msg() string {
	d.RLock()
	defer d.RUnlock()
	return d.msg
}

func (d *deployResult) Model() *types.ProjectModel {
	d.RLock()
	defer d.RUnlock()
	return d.model
}

func (d *deployResult) ResultType() websocket_pb.ResultType {
	d.RLock()
	defer d.RUnlock()
	return d.result
}

func (d *deployResult) Set(t websocket_pb.ResultType, msg string, model *types.ProjectModel) {
	d.Lock()
	defer d.Unlock()
	d.result = t
	d.msg = msg
	d.model = model
	d.set = true
}

type vars map[string]string

func (v vars) ToKeyValue() (res []*types.KeyValue) {
	for k, va := range v {
		res = append(res, &types.KeyValue{
			Key:   k,
			Value: va,
		})
	}
	return
}

func (v vars) MustGetString(key string) string {
	if value, ok := v[key]; ok {
		return value
	}

	return ""
}

func (v vars) Add(key, value string) {
	v[key] = value
}

type pipelineVars struct {
	Pipeline string
	Commit   string
	Branch   string
}

var matchTag = regexp.MustCompile(`image:\s+(\S+)`)

func matchDockerImage(v pipelineVars, manifest string) []string {
	var (
		candidateImages = make([]string, 0)
		all             = make([]string, 0)
		existsMap       = make(map[string]struct{})
	)
	submatch := matchTag.FindAllStringSubmatch(manifest, -1)
	for _, matches := range submatch {
		if len(matches) == 2 {
			image := strings.Trim(matches[1], "\"")

			if _, ok := existsMap[image]; ok {
				continue
			}
			existsMap[image] = struct{}{}
			all = append(all, image)
			if imageUsedPipelineVars(v, image) {
				candidateImages = append(candidateImages, image)
			}
		}
	}
	// 如果找到至少一个镜像就直接返回，如果未找到，则返回所有匹配到的镜像
	if len(candidateImages) > 0 {
		return candidateImages
	}

	return all
}

// imageUsedPipelineVars 使用的流水线变量的镜像，都把他当成是我们的目标镜像
func imageUsedPipelineVars(v pipelineVars, s string) bool {
	var pipelineVarsSlice []string
	if v.Pipeline != "" {
		pipelineVarsSlice = append(pipelineVarsSlice, v.Pipeline)
	}
	if v.Commit != "" {
		pipelineVarsSlice = append(pipelineVarsSlice, v.Commit)
	}
	if v.Branch != "" {
		pipelineVarsSlice = append(pipelineVarsSlice, v.Branch)
	}
	for _, pvar := range pipelineVarsSlice {
		if strings.Contains(s, pvar) {
			return true
		}
	}

	return false
}

type internalCloser func() error

func (fn internalCloser) Close() error {
	return fn()
}

func NewCloser(fn func() error) io.Closer {
	return internalCloser(fn)
}

type commit struct {
	ID             string     `json:"id"`
	ShortID        string     `json:"short_id"`
	Title          string     `json:"title"`
	CommittedDate  *time.Time `json:"committed_date"`
	AuthorName     string     `json:"author_name"`
	AuthorEmail    string     `json:"author_email"`
	CommitterName  string     `json:"committer_name"`
	CommitterEmail string     `json:"committer_email"`
	CreatedAt      *time.Time `json:"created_at"`
	Message        string     `json:"message"`
	ProjectID      int64      `json:"project_id"`
	WebURL         string     `json:"web_url"`
}

func NewEmptyCommit() application.Commit {
	return &commit{}
}

func (c *commit) GetID() string {
	return c.ID
}

func (c *commit) GetShortID() string {
	return c.ShortID
}

func (c *commit) GetTitle() string {
	return c.Title
}

func (c *commit) GetCommittedDate() *time.Time {
	return c.CommittedDate
}

func (c *commit) GetAuthorName() string {
	return c.AuthorName
}

func (c *commit) GetAuthorEmail() string {
	return c.AuthorEmail
}

func (c *commit) GetCommitterName() string {
	return c.CommitterName
}

func (c *commit) GetCommitterEmail() string {
	return c.CommitterEmail
}

func (c *commit) GetCreatedAt() *time.Time {
	return c.CreatedAt
}

func (c *commit) GetMessage() string {
	return c.Message
}

func (c *commit) GetProjectID() int64 {
	return c.ProjectID
}

func (c *commit) GetWebURL() string {
	return c.WebURL
}

type emptyPubSub struct{}

func NewEmptyPubSub() application.PubSub {
	return &emptyPubSub{}
}

func (e *emptyPubSub) Join(projectID int64) error {
	return nil
}

func (e *emptyPubSub) Leave(nsID int64, projectID int64) error {
	return nil
}

func (e *emptyPubSub) Run(ctx context.Context) error {
	return nil
}

func (e *emptyPubSub) Publish(nsID int64, pod *corev1.Pod) error {
	return nil
}

func (e *emptyPubSub) Info() any {
	return nil
}

func (e *emptyPubSub) Uid() string {
	return ""
}

func (e *emptyPubSub) ID() string {
	return ""
}

func (e *emptyPubSub) ToSelf(message application.WebsocketMessage) error {
	return nil
}

func (e *emptyPubSub) ToAll(message application.WebsocketMessage) error {
	return nil
}

func (e *emptyPubSub) ToOthers(message application.WebsocketMessage) error {
	return nil
}

func (e *emptyPubSub) Subscribe() <-chan []byte {
	return nil
}

func (e *emptyPubSub) Close() error {
	return nil
}
