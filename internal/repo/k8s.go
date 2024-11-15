package repo

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/duc-cnzj/mars/v5/internal/data"
	"github.com/duc-cnzj/mars/v5/internal/ent/schema/schematype"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/uploader"
	"github.com/duc-cnzj/mars/v5/internal/util/k8s"
	"github.com/duc-cnzj/mars/v5/internal/util/rand"
	"github.com/duc-cnzj/mars/v5/internal/util/timer"
	"github.com/dustin/go-humanize"
	"github.com/mholt/archiver/v3"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"helm.sh/helm/v3/pkg/releaseutil"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	eventv1 "k8s.io/api/events/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

const defaultContainerAnnotationName = "kubectl.kubernetes.io/default-container"

type ClusterStatus = string

const (
	StatusBad     ClusterStatus = "bad"
	StatusNotGood ClusterStatus = "not good"
	StatusHealth  ClusterStatus = "health"
)

type Container struct {
	Namespace string `json:"namespace"`
	Pod       string `json:"pod"`
	Container string `json:"container"`
}

type DockerConfig map[string]DockerConfigEntry

type DockerConfigEntry struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Auth     string `json:"auth,omitempty"`
}

type DockerConfigJSON struct {
	Auths       DockerConfig      `json:"auths"`
	HttpHeaders map[string]string `json:"HttpHeaders,omitempty"`
}

type K8sRepo interface {
	SplitManifests(manifest string) []string
	AddTlsSecret(ns string, name string, key string, crt string) (*corev1.Secret, error)
	GetPodMetrics(ctx context.Context, namespace, podName string) (*v1beta1.PodMetrics, error)
	CreateDockerSecret(ctx context.Context, namespace string) (*corev1.Secret, error)
	GetNamespace(ctx context.Context, name string) (*corev1.Namespace, error)
	CreateNamespace(ctx context.Context, name string) (*corev1.Namespace, error)
	LogStream(ctx context.Context, namespace, pod, container string) (chan []byte, error)
	GetPodLogs(ctx context.Context, namespace, podName string, options *corev1.PodLogOptions) (string, error)
	FindDefaultContainer(ctx context.Context, namespace string, pod string) (string, error)
	GetPod(namespace, podName string) (*corev1.Pod, error)
	ListEvents(namespace string) ([]*eventv1.Event, error)
	IsPodRunning(namespace, podName string) (running bool, notRunningReason string)
	GetPodSelectorsByManifest(manifests []string) []string
	GetCpuAndMemoryInNamespace(ctx context.Context, namespace string) (string, string)
	GetCpuAndMemory(ctx context.Context, list []v1beta1.PodMetrics) (string, string)
	GetCpuAndMemoryQuantity(pod v1beta1.PodMetrics) (cpu *resource.Quantity, memory *resource.Quantity)
	ClusterInfo() *ClusterInfo

	Execute(ctx context.Context, c *Container, input *ExecuteInput) error
	DeleteSecret(ctx context.Context, namespace, secret string) error
	DeleteNamespace(ctx context.Context, name string) error

	GetAllPodMetrics(ctx context.Context, proj *Project) []v1beta1.PodMetrics

	CopyFileToPod(ctx context.Context, input *CopyFileToPodInput) (*File, error)
	CopyFromPod(ctx context.Context, input *CopyFromPodInput) (*File, error)
}

var _ K8sRepo = (*k8sRepo)(nil)

type k8sRepo struct {
	logger        mlog.Logger
	uploader      uploader.Uploader
	maxUploadSize uint64
	archiver      Archiver
	executor      ExecutorManager
	data          data.Data
	fileRepo      FileRepo
	timer         timer.Timer
}

func NewK8sRepo(
	logger mlog.Logger,
	timer timer.Timer,
	data data.Data,
	fileRepo FileRepo,
	uploader uploader.Uploader,
	archiver Archiver,
	remoteExecutor ExecutorManager,
) K8sRepo {
	return &k8sRepo{
		fileRepo:      fileRepo,
		timer:         timer,
		data:          data,
		logger:        logger.WithModule("repo/k8s"),
		uploader:      uploader,
		maxUploadSize: data.Config().MaxUploadSize(),
		archiver:      archiver,
		executor:      remoteExecutor,
	}
}

type CopyFromPodInput struct {
	Namespace string
	Pod       string
	Container string
	// pod 内的绝对路径
	FilePath string
	UserName string
}

func (repo *k8sRepo) CopyFromPod(ctx context.Context, input *CopyFromPodInput) (*File, error) {
	ctx, span := otel.Tracer("").Start(ctx, "CopyFromPod")
	defer span.End()

	var (
		file uploader.File
		err  error
	)

	lsbf := &bytes.Buffer{}

	_ = repo.Execute(ctx, &Container{
		Namespace: input.Namespace,
		Pod:       input.Pod,
		Container: input.Container,
	}, &ExecuteInput{
		Stdout: lsbf,
		TTY:    false,
		Cmd:    []string{"sh", "-c", fmt.Sprintf("test -f %s && echo 1 || echo 0", input.FilePath)},
	})
	isFile := strings.Trim(lsbf.String(), "\n") == "1"
	if !isFile {
		return nil, ToError(400, "下载内容必须是文件")
	}

	pwdbf := &bytes.Buffer{}
	err = repo.Execute(ctx, &Container{
		Namespace: input.Namespace,
		Pod:       input.Pod,
		Container: input.Container,
	}, &ExecuteInput{
		Stdout: pwdbf,
		TTY:    false,
		Cmd:    []string{"sh", "-c", "pwd"},
	})
	if err != nil {
		return nil, err
	}
	base := strings.Trim(pwdbf.String(), "\n")

	if !strings.HasPrefix(input.FilePath, base) {
		return nil, ToError(400, "invalid file path")
	}

	rel, err := filepath.Rel(base, input.FilePath)
	if err != nil {
		return nil, err
	}
	input.FilePath = rel
	repo.logger.Debugf("[CopyFromPod]: rel: %q base: %q", rel, base)

	bf := &bytes.Buffer{}
	fileCopy := repo.executor.NewFileCopy(5, bf)

	remotePath := input.FilePath

	filename := fmt.Sprintf("%s-%s.tar", input.Pod, rand.String(10))

	up := repo.uploader.Disk("podfile")

	file, err = up.NewFile(
		fmt.Sprintf("%s/%s/%s/%s",
			input.UserName,
			repo.timer.Now().Format("2006-01-02"),
			fmt.Sprintf("%s-%s", repo.timer.Now().Format("15-04-05"), rand.String(20)),
			filename),
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		file.Close()
		if err != nil {
			up.Delete(file.Name())
		}
	}()

	if err = fileCopy.CopyFromPod(ctx, k8s.CopyFileSpec{
		PodName:       input.Pod,
		PodNamespace:  input.Namespace,
		ContainerName: input.Container,
		File:          k8s.NewRemotePath(remotePath),
	}, file); err != nil {
		return nil, fmt.Errorf("copy from pod error: %w, %v", err, bf.String())
	}

	var stat os.FileInfo
	stat, err = file.Stat()
	if err != nil {
		repo.logger.Error(err)
		return nil, err
	}
	return repo.fileRepo.Create(ctx, &CreateFileInput{
		Path:       file.Name(),
		Username:   input.UserName,
		Size:       uint64(stat.Size()),
		UploadType: up.Type(),
		Namespace:  input.Namespace,
		Pod:        input.Pod,
		Container:  input.Container,
	})
}

// SplitManifests
// 因为有些 secret 自带 --- 的值，导致 spilt "---" 解析异常
func (repo *k8sRepo) SplitManifests(manifest string) []string {
	mapManifests := releaseutil.SplitManifests(manifest)
	var manifests = make([]string, 0, len(mapManifests))
	for _, s := range mapManifests {
		manifests = append(manifests, s)
	}
	sort.Strings(manifests)

	return manifests
}

func (repo *k8sRepo) AddTlsSecret(ns string, name string, key string, crt string) (*corev1.Secret, error) {
	return repo.data.K8sClient().Client.CoreV1().Secrets(ns).Create(context.TODO(), &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Annotations: map[string]string{
				"created-by": "mars",
			},
		},
		StringData: map[string]string{
			"tls.key": key,
			"tls.crt": crt,
		},
		Type: corev1.SecretTypeTLS,
	}, metav1.CreateOptions{})
}

func (repo *k8sRepo) GetPodMetrics(ctx context.Context, namespace, podName string) (*v1beta1.PodMetrics, error) {
	return repo.data.K8sClient().MetricsClient.MetricsV1beta1().PodMetricses(namespace).Get(ctx, podName, metav1.GetOptions{})
}

func (repo *k8sRepo) GetAllPodMetrics(ctx context.Context, proj *Project) []v1beta1.PodMetrics {
	_, span := otel.Tracer("").Start(ctx, "GetAllPodMetrics")
	defer span.End()
	metricses := repo.data.K8sClient().MetricsClient.MetricsV1beta1().PodMetricses(proj.Namespace.Name)
	var list []v1beta1.PodMetrics
	if len(proj.PodSelectors) == 0 {
		return nil
	}
	for _, podlabels := range proj.PodSelectors {
		labelsMap, _ := labels.ConvertSelectorToLabelsMap(podlabels)
		ret, _ := repo.data.K8sClient().PodLister.Pods(proj.Namespace.Name).List(labelsMap.AsSelector())
		if len(ret) == 0 {
			continue
		}
		l, _ := metricses.List(ctx, metav1.ListOptions{
			LabelSelector: podlabels,
		})
		repo.logger.DebugCtx(ctx, "[GetAllPodMetrics]: ", podlabels, " ", len(l.Items))

		list = append(list, l.Items...)
	}

	return list
}

func (repo *k8sRepo) DeleteNamespace(ctx context.Context, name string) error {
	return repo.data.K8sClient().Client.CoreV1().Namespaces().Delete(ctx, name, metav1.DeleteOptions{})
}

func (repo *k8sRepo) DeleteSecret(ctx context.Context, namespace, secret string) error {
	return repo.data.K8sClient().Client.CoreV1().Secrets(namespace).Delete(ctx, secret, metav1.DeleteOptions{})
}

func (repo *k8sRepo) CreateDockerSecret(ctx context.Context, namespace string) (*corev1.Secret, error) {
	var entries = make(map[string]DockerConfigEntry)
	for _, auth := range repo.data.Config().ImagePullSecrets {
		entries[auth.Server] = DockerConfigEntry{
			Username: auth.Username,
			Password: auth.Password,
			Email:    auth.Email,
			Auth:     base64.StdEncoding.EncodeToString([]byte(auth.Username + ":" + auth.Password)),
		}
	}

	dockerCfgJSON := DockerConfigJSON{
		Auths: entries,
	}

	marshal, _ := json.Marshal(dockerCfgJSON)

	return repo.data.K8sClient().Client.CoreV1().Secrets(namespace).Create(ctx, &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: corev1.SchemeGroupVersion.String(),
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      "mars-" + strings.ToLower(rand.String(10)),
		},
		Data: map[string][]byte{
			corev1.DockerConfigJsonKey: marshal,
		},
		Type: corev1.SecretTypeDockerConfigJson,
	}, metav1.CreateOptions{})
}

func (repo *k8sRepo) GetNamespace(ctx context.Context, name string) (*corev1.Namespace, error) {
	return repo.data.K8sClient().Client.CoreV1().Namespaces().Get(ctx, name, metav1.GetOptions{})
}

func (repo *k8sRepo) CreateNamespace(ctx context.Context, name string) (*corev1.Namespace, error) {
	return repo.data.K8sClient().Client.CoreV1().
		Namespaces().
		Create(ctx,
			&corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
				},
			},
			metav1.CreateOptions{},
		)
}

func (repo *k8sRepo) Execute(ctx context.Context, c *Container, input *ExecuteInput) error {
	return repo.executor.New().
		WithContainer(c.Namespace, c.Pod, c.Container).
		WithMethod("POST").
		WithCommand(input.Cmd).
		Execute(ctx, input)
}

func (repo *k8sRepo) GetPodLogs(ctx context.Context, namespace, podName string, options *corev1.PodLogOptions) (string, error) {
	logs := repo.data.K8sClient().Client.CoreV1().Pods(namespace).GetLogs(podName, options)
	do := logs.Do(ctx)
	raw, err := do.Raw()
	return string(raw), err
}

func (repo *k8sRepo) ListEvents(namespace string) ([]*eventv1.Event, error) {
	return repo.data.K8sClient().EventLister.Events(namespace).List(labels.Everything())
}

func (repo *k8sRepo) FindDefaultContainer(ctx context.Context, namespace string, pod string) (string, error) {
	corev1pod, err := repo.GetPod(namespace, pod)
	if err != nil {
		return "", err
	}
	if name := corev1pod.Annotations[defaultContainerAnnotationName]; len(name) > 0 {
		for _, co := range corev1pod.Spec.Containers {
			if name == co.Name {
				return name, nil
			}
		}
	}

	for _, co := range corev1pod.Spec.Containers {
		return co.Name, nil
	}

	return "", errors.New("no container found")
}

func (repo *k8sRepo) GetPod(namespace, podName string) (*corev1.Pod, error) {
	return repo.data.K8sClient().PodLister.Pods(namespace).Get(podName)
}

func (repo *k8sRepo) IsPodRunning(namespace, podName string) (running bool, notRunningReason string) {
	podInfo, err := repo.data.K8sClient().PodLister.Pods(namespace).Get(podName)
	if err != nil {
		return false, err.Error()
	}

	if podInfo.Status.Phase == corev1.PodRunning {
		return true, ""
	}

	if podInfo.Status.Phase == corev1.PodFailed && podInfo.Status.Reason == "Evicted" {
		return false, fmt.Sprintf("po %s already evicted in namespace %s!", podName, namespace)
	}

	for _, status := range podInfo.Status.ContainerStatuses {
		return false, fmt.Sprintf("%s %s", status.State.Waiting.Reason, status.State.Waiting.Message)
	}

	return false, "pod not running."
}

// GetPodSelectorsByManifest
// 参考 https://github.com/kubernetes/client-go/issues/193#issuecomment-363240636
// 参考源码
func (repo *k8sRepo) GetPodSelectorsByManifest(manifests []string) []string {
	var selectors []string
	info, _ := runtime.SerializerInfoForMediaType(scheme.Codecs.SupportedMediaTypes(), runtime.ContentTypeYAML)
	for _, f := range manifests {
		obj, _, _ := info.Serializer.Decode([]byte(f), nil, nil)
		switch a := obj.(type) {
		case *appsv1.Deployment:
			selector, _ := metav1.LabelSelectorAsSelector(a.Spec.Selector)
			selectors = append(selectors, selector.String())
		case *appsv1.StatefulSet:
			selector, _ := metav1.LabelSelectorAsSelector(a.Spec.Selector)
			selectors = append(selectors, selector.String())
		case *appsv1.DaemonSet:
			selector, _ := metav1.LabelSelectorAsSelector(a.Spec.Selector)
			selectors = append(selectors, selector.String())
		case *batchv1.Job:
			jobPodLabels := a.Spec.Template.Labels
			if jobPodLabels != nil {
				selectors = append(selectors, labels.SelectorFromSet(jobPodLabels).String())
			}
		case *batchv1beta1.CronJob:
			jobPodLabels := a.Spec.JobTemplate.Spec.Template.Labels
			if jobPodLabels != nil {
				selectors = append(selectors, labels.SelectorFromSet(jobPodLabels).String())
			}
		case *batchv1.CronJob:
			jobPodLabels := a.Spec.JobTemplate.Spec.Template.Labels
			if jobPodLabels != nil {
				selectors = append(selectors, labels.SelectorFromSet(jobPodLabels).String())
			}
		default:
			repo.logger.Debugf("GetPodSelectorsByManifest Default: %#v", a)
		}
	}

	return selectors
}

func (repo *k8sRepo) GetCpuAndMemoryInNamespace(ctx context.Context, namespace string) (string, string) {
	metricses := repo.data.K8sClient().MetricsClient.MetricsV1beta1().PodMetricses(namespace)
	list, _ := metricses.List(ctx, metav1.ListOptions{})
	return repo.GetCpuAndMemory(ctx, list.Items)
}

func (repo *k8sRepo) GetCpuAndMemory(ctx context.Context, list []v1beta1.PodMetrics) (string, string) {
	_, span := otel.Tracer("").Start(ctx, "GetCpuAndMemory")
	defer span.End()
	var cpu, memory *resource.Quantity

	for _, item := range list {
		for _, container := range item.Containers {
			if cpu == nil {
				cpu = container.Usage.Cpu()
			} else {
				cpu.Add(*container.Usage.Cpu())
			}

			if memory == nil {
				memory = container.Usage.Memory()
			} else {
				memory.Add(*container.Usage.Memory())
			}
		}
	}

	var cpuStr, memoryStr string = "0 m", "0 MB"

	if cpu != nil {
		cpuStr = fmt.Sprintf("%d m", cpu.MilliValue())
	}
	if memory != nil {
		asInt64, _ := memory.AsInt64()
		memoryStr = humanize.Bytes(uint64(asInt64))
	}

	return cpuStr, memoryStr
}

func (repo *k8sRepo) GetCpuAndMemoryQuantity(pod v1beta1.PodMetrics) (cpu *resource.Quantity, memory *resource.Quantity) {
	for _, container := range pod.Containers {
		if cpu == nil {
			cpu = container.Usage.Cpu()
		} else {
			cpu.Add(*container.Usage.Cpu())
		}

		if memory == nil {
			memory = container.Usage.Memory()
		} else {
			memory.Add(*container.Usage.Memory())
		}
	}

	return cpu, memory
}

type copyToPodResult struct {
	TargetDir     string
	ErrOut        string
	StdOut        string
	ContainerPath string
	FileName      string
}

func (repo *k8sRepo) copyToPod(ctx context.Context, namespace, pod, container, fpath, targetContainerDir string) (*copyToPodResult, error) {
	tracer := otel.Tracer("")
	ctx, span := tracer.Start(ctx, "k8sRepo/copyToPod")
	defer span.End()
	span.SetAttributes(
		attribute.Key("namespace").String(namespace),
		attribute.Key("pod").String(pod),
		attribute.Key("container").String(container),
		attribute.Key("file").String(fpath),
	)

	var (
		errbf, outbf      = bytes.NewBuffer([]byte{}), bytes.NewBuffer([]byte{})
		reader, outStream = io.Pipe()
		uploader          = repo.uploader
		localUploader     = repo.uploader.LocalUploader()
	)
	if targetContainerDir == "" {
		targetContainerDir = "/tmp"
	}
	st, err := uploader.Stat(fpath)
	if err != nil {
		return nil, err
	}
	if st.Size() > repo.maxUploadSize {
		return nil, fmt.Errorf("最大不得超过 %s, 你上传的文件大小是 %s", humanize.Bytes(repo.maxUploadSize), humanize.Bytes(uint64(st.Size())))
	}

	baseName := filepath.Base(fpath)
	path := filepath.Join(filepath.Dir(fpath), baseName+".tar.gz")
	repo.logger.Debugf("[CopyFileToPod]: %v", path)
	var localPath string = fpath
	// 如果是非 local 类型的，需要远程下载到 local 进行打包，再上传到容器
	if uploader.Type() != schematype.Local {
		read, err := uploader.Read(fpath)
		if err != nil {
			return nil, err
		}
		defer read.Close()
		if localUploader.Exists(localPath) {
			localUploader.Delete(localPath)
		}
		put, err := localUploader.Put(localPath, read)
		if err != nil {
			return nil, err
		}
		localPath = put.Path()
		defer localUploader.Delete(localPath)
	}
	if err := repo.archiver.Archive([]string{localPath}, path); err != nil {
		return nil, err
	}
	defer repo.archiver.Remove(path)
	src, err := repo.archiver.Open(path)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()
	go func(reader *io.PipeReader, outStream *io.PipeWriter, src io.ReadCloser) {
		defer func() {
			reader.Close()
			outStream.Close()
			src.Close()
			wg.Done()
		}()
		defer repo.logger.HandlePanic("CopyFileToPod")

		if _, err := io.Copy(outStream, src); err != nil {
			repo.logger.Error(err)
		}
	}(reader, outStream, src)

	err = repo.executor.
		New().
		WithCommand([]string{"tar", "-zmxf", "-", "-C", targetContainerDir}).
		WithMethod("POST").
		WithContainer(namespace, pod, container).
		Execute(
			ctx,
			&ExecuteInput{
				Stdin:  reader,
				Stdout: outbf,
				Stderr: errbf,
				TTY:    false,
			},
		)

	return &copyToPodResult{
		TargetDir:     targetContainerDir,
		ErrOut:        errbf.String(),
		StdOut:        outbf.String(),
		ContainerPath: filepath.Join(targetContainerDir, baseName),
		FileName:      baseName,
	}, err
}

type CopyFileToPodInput struct {
	FileId    int64
	Namespace string
	Pod       string
	Container string
}

func (repo *k8sRepo) CopyFileToPod(ctx context.Context, input *CopyFileToPodInput) (*File, error) {
	file, err := repo.fileRepo.GetByID(ctx, int(input.FileId))
	if err != nil {
		repo.logger.ErrorCtx(ctx, err)
		return nil, err
	}

	result, err := repo.copyToPod(ctx, input.Namespace, input.Pod, input.Container, file.Path, "")
	if err != nil {
		repo.logger.ErrorCtx(ctx, err)
		return nil, err
	}

	return repo.fileRepo.Update(ctx, &UpdateFileRequest{
		ID:            int(input.FileId),
		ContainerPath: result.ContainerPath,
		Namespace:     input.Namespace,
		Pod:           input.Pod,
		Container:     input.Container,
	})
}

type ClusterInfo struct {
	// 健康状况
	Status ClusterStatus `json:"status"`

	// 可用内存
	FreeMemory string `json:"free_memory"`
	// 可用 cpu
	FreeCpu string `json:"free_cpu"`

	// 可分配内存
	FreeRequestMemory string `json:"free_request_memory"`
	// 可分配 cpu
	FreeRequestCpu string `json:"free_request_cpu"`

	// 总共的可调度的内存
	TotalMemory string `json:"total_memory"`
	// 总共的可调度的 cpu
	TotalCpu string `json:"total_cpu"`

	// 内存使用率
	UsageMemoryRate string `json:"usage_memory_rate"`
	// cpu 使用率
	UsageCpuRate string `json:"usage_cpu_rate"`

	// 内存分配率
	RequestMemoryRate string `json:"request_memory_rate"`
	// cpu 分配率
	RequestCpuRate string `json:"request_cpu_rate"`
}

func (repo *k8sRepo) ClusterInfo() *ClusterInfo {
	selector := labels.Everything()
	var nodes []corev1.Node

	// 获取已经使用的 cpu, memory
	nodeList, _ := repo.data.K8sClient().Client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
		LabelSelector: selector.String(),
	})
	nodes = append(nodes, nodeList.Items...)
	allocatable := make(map[string]corev1.ResourceList)

	var (
		totalCpu    = &resource.Quantity{}
		totalMemory = &resource.Quantity{}
	)

	var (
		workerNodes  []corev1.Node
		notWorkNodes []corev1.Node
	)

	for _, node := range nodes {
		notWork := false
		for _, taint := range node.Spec.Taints {
			if taint.Effect == corev1.TaintEffectNoExecute || taint.Effect == corev1.TaintEffectNoSchedule {
				notWork = true
				break
			}
		}
		if !notWork {
			workerNodes = append(workerNodes, node)
		} else {
			notWorkNodes = append(notWorkNodes, node)
		}
	}

	for _, n := range workerNodes {
		allocatable[n.Name] = n.Status.Allocatable
		totalCpu.Add(n.Status.Allocatable.Cpu().DeepCopy())
		totalMemory.Add(n.Status.Allocatable.Memory().DeepCopy())
	}

	requestCpu, requestMemory := repo.getNodeRequestCpuAndMemory(notWorkNodes)
	var (
		usedCpu    = &resource.Quantity{}
		usedMemory = &resource.Quantity{}
	)

	list, _ := repo.data.K8sClient().MetricsClient.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{
		LabelSelector: selector.String(),
	})

	IsStatisticalNode := func(workerNodes []corev1.Node, name string) bool {
		for _, node := range workerNodes {
			if node.Name == name {
				return true
			}
		}
		return false
	}

	for _, item := range list.Items {
		if !IsStatisticalNode(workerNodes, item.Name) {
			continue
		}
		usedCpu.Add(item.Usage.Cpu().DeepCopy())
		usedMemory.Add(item.Usage.Memory().DeepCopy())
	}

	freeMemory := totalMemory.DeepCopy()
	freeMemory.Sub(*usedMemory)
	freeCpu := totalCpu.DeepCopy()
	freeCpu.Sub(*usedCpu)

	freeRequestMemory := totalMemory.DeepCopy()
	freeRequestMemory.Sub(*requestMemory)
	freeRequestCpu := totalCpu.DeepCopy()
	freeRequestCpu.Sub(*requestCpu)

	rateMemory := float64(usedMemory.Value()) / float64(totalMemory.Value()) * 100
	rateCpu := float64(usedCpu.Value()) / float64(totalCpu.Value()) * 100
	rateRequestMemory := float64(requestMemory.Value()) / float64(totalMemory.Value()) * 100
	rateRequestCpu := float64(requestCpu.Value()) / float64(totalCpu.Value()) * 100

	return &ClusterInfo{
		Status:            repo.getStatus(rateRequestMemory, rateRequestCpu),
		FreeRequestMemory: humanize.IBytes(uint64(freeRequestMemory.Value())),
		FreeRequestCpu:    fmt.Sprintf("%.2f core", float64(freeRequestCpu.MilliValue())/1000),
		FreeMemory:        humanize.IBytes(uint64(freeMemory.Value())),
		FreeCpu:           fmt.Sprintf("%.2f core", float64(freeCpu.MilliValue())/1000),
		TotalMemory:       humanize.IBytes(uint64(totalMemory.Value())),
		TotalCpu:          fmt.Sprintf("%.2f core", float64(totalCpu.MilliValue())/1000),
		UsageMemoryRate:   fmt.Sprintf("%.1f%%", rateMemory),
		UsageCpuRate:      fmt.Sprintf("%.1f%%", rateCpu),
		RequestCpuRate:    fmt.Sprintf("%.1f%%", rateRequestCpu),
		RequestMemoryRate: fmt.Sprintf("%.1f%%", rateRequestMemory),
	}
}

func (repo *k8sRepo) getStatus(rateRequestMemory float64, rateRequestCpu float64) ClusterStatus {
	var status = StatusHealth
	if rateRequestMemory > 60 || rateRequestCpu > 60 {
		status = StatusNotGood
	}
	if rateRequestMemory > 80 || rateRequestCpu > 80 {
		status = StatusBad
	}
	return status
}

func (repo *k8sRepo) getNodeRequestCpuAndMemory(noExecuteNodes []corev1.Node) (*resource.Quantity, *resource.Quantity) {
	var (
		requestCpu    = &resource.Quantity{}
		requestMemory = &resource.Quantity{}
	)

	var nodeSelector []string = []string{
		"status.phase==" + string(corev1.PodRunning),
	}
	for _, node := range noExecuteNodes {
		nodeSelector = append(nodeSelector, "spec.nodeName!="+node.Name)
	}
	fieldSelector, _ := fields.ParseSelector(strings.Join(nodeSelector, ","))
	nodeNonTerminatedPodsList, _ := repo.data.K8sClient().
		Client.
		CoreV1().
		Pods("").
		List(context.TODO(), metav1.ListOptions{FieldSelector: fieldSelector.String()})
	for _, item := range nodeNonTerminatedPodsList.Items {
		for _, container := range item.Spec.Containers {
			requestCpu.Add(container.Resources.Requests.Cpu().DeepCopy())
			requestMemory.Add(container.Resources.Requests.Memory().DeepCopy())
		}
	}

	return requestCpu, requestMemory
}

func (repo *k8sRepo) LogStream(
	ctx context.Context,
	namespace,
	pod,
	container string,
) (chan []byte, error) {
	tailLines := 1000
	logs := repo.data.K8sClient().Client.
		CoreV1().
		Pods(namespace).
		GetLogs(pod, &corev1.PodLogOptions{
			Follow:    true,
			Container: container,
			TailLines: lo.ToPtr(int64(tailLines)),
		})

	stream, err := logs.Stream(ctx)
	if err != nil {
		return nil, err
	}

	bf := bufio.NewReader(stream)
	ctx, cancelFunc := context.WithCancel(ctx)
	ch := make(chan []byte, tailLines)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			close(ch)
			cancelFunc()
			wg.Done()
		}()
		defer repo.logger.HandlePanic("StreamContainerLog")

		for {
			bytes, err := bf.ReadBytes('\n')
			if err != nil {
				repo.logger.DebugCtxf(ctx, "[LogStream]: %v", err)
				return
			}
			select {
			case ch <- bytes:
			default:
				repo.logger.DebugCtx(ctx, "[LogStream]:  drop line!")
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		stream.Close()
		repo.logger.DebugCtx(ctx, "[LogStream]:  ctx done!")
	}()
	go func() {
		wg.Wait()
		repo.logger.DebugCtx(ctx, "[LogStream]:  exit!")
	}()

	return ch, nil
}

type Archiver interface {
	Archive(sources []string, destination string) error
	Open(path string) (io.ReadCloser, error)
	Remove(path string) error
}

type defaultArchiver struct{}

func NewDefaultArchiver() Archiver {
	return &defaultArchiver{}
}

func (m *defaultArchiver) Archive(sources []string, destination string) error {
	return archiver.Archive(sources, destination)
}

func (m *defaultArchiver) Open(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func (m *defaultArchiver) Remove(path string) error {
	return os.Remove(path)
}

type ExecutorManager interface {
	New() Executor
	NewFileCopy(maxTries int, errOut io.Writer) k8s.FileCopy
}

type ExecuteInput struct {
	Stdin             io.Reader
	Stdout, Stderr    io.Writer
	TTY               bool
	Cmd               []string
	TerminalSizeQueue remotecommand.TerminalSizeQueue
}

type Executor interface {
	WithMethod(method string) Executor
	WithContainer(namespace, pod, container string) Executor
	WithCommand(cmd []string) Executor
	Execute(context.Context, *ExecuteInput) error
}

type defaultRemoteExecutor struct {
	data   data.Data
	logger mlog.Logger
}

func NewExecutorManager(data data.Data, logger mlog.Logger) ExecutorManager {
	return &defaultRemoteExecutor{
		data:   data,
		logger: logger.WithModule("ExecutorManager"),
	}
}

func (e *defaultRemoteExecutor) NewFileCopy(
	maxTries int,
	errOut io.Writer,
) k8s.FileCopy {
	restCfg := e.data.K8sClient().RestConfig
	restCfg.GroupVersion = &schema.GroupVersion{Group: "", Version: "v1"}
	restCfg.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	restCfg.APIPath = "/api"

	return k8s.NewCopyOptions(
		e.logger,
		restCfg,
		e.data.K8sClient().Client,
		maxTries,
		errOut,
	)
}

func (e *defaultRemoteExecutor) New() Executor {
	return &executor{
		clientSet: e.data.K8sClient().Client,
		config:    e.data.K8sClient().RestConfig,
	}
}

type executor struct {
	namespace, pod, container string
	method                    string
	cmd                       []string

	clientSet kubernetes.Interface
	config    *restclient.Config
}

func (e *executor) WithMethod(method string) Executor {
	e.method = method
	return e
}

func (e *executor) WithContainer(namespace, pod, container string) Executor {
	e.namespace = namespace
	e.pod = pod
	e.container = container
	return e
}

func (e *executor) WithCommand(cmd []string) Executor {
	e.cmd = cmd
	return e
}

func (e *executor) Execute(ctx context.Context, input *ExecuteInput) error {
	var (
		terminalSizeQueue = input.TerminalSizeQueue
		stdin             = input.Stdin
		stdout            = input.Stdout
		stderr            = input.Stderr
		tty               = input.TTY
	)
	peo := e.newOption(stdin, stdout, stderr, tty)

	req := e.clientSet.CoreV1().
		RESTClient().
		Post().
		Namespace(e.namespace).
		Resource("pods").
		SubResource("exec").
		Name(e.pod)

	exec, err := remotecommand.NewSPDYExecutor(e.config, e.method, req.VersionedParams(peo, scheme.ParameterCodec).URL())
	if err != nil {
		return err
	}

	return exec.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdin:             stdin,
		Stdout:            stdout,
		Stderr:            stderr,
		Tty:               tty,
		TerminalSizeQueue: terminalSizeQueue,
	})
}

func (e *executor) newOption(stdin io.Reader, stdout io.Writer, stderr io.Writer, tty bool) *corev1.PodExecOptions {
	return &corev1.PodExecOptions{
		Stdin:     stdin != nil,
		Stdout:    stdout != nil,
		Stderr:    stderr != nil,
		TTY:       tty,
		Container: e.container,
		Command:   e.cmd,
	}
}
