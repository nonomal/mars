package services

import (
	"context"
	"errors"
	"io"
	"slices"
	"sort"
	"testing"
	"time"

	"github.com/duc-cnzj/mars/api/v5/container"
	"github.com/duc-cnzj/mars/api/v5/types"
	"github.com/duc-cnzj/mars/v5/internal/auth"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/repo"
	"github.com/duc-cnzj/mars/v5/internal/util/timer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	eventv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/remotecommand"
	clientgoexec "k8s.io/client-go/util/exec"
)

func TestNewContainerSvc(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		repo.NewMockK8sRepo(m),
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	assert.NotNil(t, svc)
	assert.NotNil(t, svc.(*containerSvc).eventRepo)
	assert.NotNil(t, svc.(*containerSvc).k8sRepo)
	assert.NotNil(t, svc.(*containerSvc).fileRepo)
	assert.NotNil(t, svc.(*containerSvc).eventRepo)
	assert.NotNil(t, svc.(*containerSvc).logger)
}

func Test_containerSvc_IsPodRunning(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(false, "")
	running, err := svc.IsPodRunning(context.TODO(), &container.IsPodRunningRequest{
		Namespace: "a",
		Pod:       "b",
	})
	assert.False(t, running.Running)
	assert.Nil(t, err)
}

func Test_containerSvc_IsPodExists(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(nil, nil)
	running, err := svc.IsPodExists(context.TODO(), &container.IsPodExistsRequest{
		Namespace: "a",
		Pod:       "b",
	})
	assert.True(t, running.Exists)
	assert.Nil(t, err)
}

func Test_containerSvc_IsPodExists_Fail(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(nil, errors.New("X"))
	running, err := svc.IsPodExists(context.TODO(), &container.IsPodExistsRequest{
		Namespace: "a",
		Pod:       "b",
	})
	assert.False(t, running.Exists)
	assert.Nil(t, err)
}

func TestContainerSvc_ContainerLog_PodNotFound(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(nil, nil)
	_, err := svc.ContainerLog(context.TODO(), &container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	})
	assert.NotNil(t, err)
}

func TestContainerSvc_ContainerLog_PodPending(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodPending}}, nil)
	_, err := svc.ContainerLog(context.TODO(), &container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	})
	assert.Error(t, err)
}

func TestContainerSvc_ContainerLog_PodRunning(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodRunning}}, nil)
	k8sRepo.EXPECT().GetPodLogs(gomock.Any(), "a", "b", &v1.PodLogOptions{
		TailLines: &tailLines,
		Container: "c",
	}).Return("log", nil)
	_, err := svc.ContainerLog(context.TODO(), &container.LogRequest{
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	})
	assert.Nil(t, err)
}

func TestContainerSvc_ContainerLog_GetPodLogs_error(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodRunning}}, nil)
	k8sRepo.EXPECT().GetPodLogs(gomock.Any(), "a", "b", gomock.Any()).Return("", errors.New("x"))
	_, err := svc.ContainerLog(context.TODO(), &container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	})
	assert.Equal(t, "x", err.Error())
}

func TestContainerSvc_ContainerLog_PodPending1(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodPending}}, nil)
	k8sRepo.EXPECT().ListEvents(gomock.Any()).Return([]*eventv1.Event{
		{
			Regarding: v1.ObjectReference{Kind: "Pod", Name: "b"},
			Note:      "aaa",
		},
		{
			Regarding: v1.ObjectReference{Kind: "Pod", Name: "b"},
			Note:      "bbb",
		},
	}, nil)
	resp, err := svc.ContainerLog(context.TODO(), &container.LogRequest{
		Namespace:  "a",
		Pod:        "b",
		ShowEvents: true,
	})
	assert.Nil(t, err)
	assert.Equal(t, "aaa\nbbb", resp.Log)
}

type logStreamServer struct {
	ctx context.Context
	container.Container_StreamContainerLogServer
	res []string
}

func (l *logStreamServer) Send(response *container.LogResponse) error {
	l.res = append(l.res, response.Log)
	return nil
}

func (l *logStreamServer) Context() context.Context {
	if l.ctx != nil {
		return l.ctx
	}
	return context.TODO()
}

func TestContainerSvc_CopyToPod_PodNotRunning(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(false, "")
	_, err := svc.CopyToPod(context.TODO(), &container.CopyToPodRequest{
		Namespace: "a",
		Pod:       "b",
	})
	assert.NotNil(t, err)
}

func TestContainerSvc_CopyToPod_Success(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	eventRepo := repo.NewMockEventRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		eventRepo,
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(true, "")
	eventRepo.EXPECT().FileAuditLog(
		types.EventActionType_Upload,
		MustGetUser(newAdminUserCtx()).Name,
		gomock.Any(),
		11,
	)
	k8sRepo.EXPECT().CopyFileToPod(gomock.Any(), &repo.CopyFileToPodInput{
		FileId:    1,
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	}).Return(&repo.File{ID: 11}, nil)
	_, err := svc.CopyToPod(newAdminUserCtx(), &container.CopyToPodRequest{
		FileId:    1,
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	})
	assert.Nil(t, err)
}
func TestContainerSvc_CopyToPod_Error(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	eventRepo := repo.NewMockEventRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		eventRepo,
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(true, "")
	k8sRepo.EXPECT().CopyFileToPod(gomock.Any(), &repo.CopyFileToPodInput{
		FileId:    1,
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	}).Return(nil, errors.New("xx"))
	_, err := svc.CopyToPod(newAdminUserCtx(), &container.CopyToPodRequest{
		FileId:    1,
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	})
	assert.Equal(t, "xx", err.Error())
}

func TestContainerSvc_StreamContainerLog_PodNotFound(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(nil, nil)
	err := svc.StreamContainerLog(&container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	}, &logStreamServer{})
	assert.NotNil(t, err)
}

func TestContainerSvc_StreamContainerLog_PodPending(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodPending}}, nil)
	err := svc.StreamContainerLog(&container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	}, &logStreamServer{})
	assert.Error(t, err)
}

func TestContainerSvc_StreamContainerLog_PodRunning(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodRunning}}, nil)
	ch := make(chan []byte, 2)
	ch <- []byte("log1")
	ch <- []byte("log2")
	close(ch)
	k8sRepo.EXPECT().LogStream(gomock.Any(), "a", "b", "c").Return(ch, nil)
	s := &logStreamServer{}
	err := svc.StreamContainerLog(&container.LogRequest{
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	}, s)
	assert.Nil(t, err)
	assert.Equal(t, []string{"log1", "log2"}, s.res)
}

func TestContainerSvc_StreamContainerLog_PodSucceeded(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodSucceeded}}, nil).AnyTimes()
	k8sRepo.EXPECT().GetPodLogs(gomock.Any(), "a", "b", gomock.Any()).Return("log", nil)
	s := &logStreamServer{}
	err := svc.StreamContainerLog(&container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	}, s)
	assert.Nil(t, err)
	assert.Equal(t, []string{"log"}, s.res)
}

func TestContainerSvc_StreamContainerLog_Error(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodSucceeded}}, nil).AnyTimes()
	k8sRepo.EXPECT().GetPodLogs(gomock.Any(), "a", "b", gomock.Any()).Return("", errors.New("x"))
	s := &logStreamServer{}
	err := svc.StreamContainerLog(&container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	}, s)
	assert.Equal(t, "x", err.Error())
}

func TestContainerSvc_StreamContainerLog_PodFailed(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodFailed}}, nil).AnyTimes()
	k8sRepo.EXPECT().GetPodLogs(gomock.Any(), "a", "b", gomock.Any()).Return("log", nil)
	s := &logStreamServer{}
	err := svc.StreamContainerLog(&container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	}, s)
	assert.Nil(t, err)
	assert.Equal(t, []string{"log"}, s.res)
}

func TestContainerSvc_StreamContainerLog_PodPending1(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().GetPod("a", "b").Return(&v1.Pod{Status: v1.PodStatus{Phase: v1.PodPending}}, nil).AnyTimes()
	s := &logStreamServer{}
	err := svc.StreamContainerLog(&container.LogRequest{
		Namespace: "a",
		Pod:       "b",
	}, s)

	assert.Equal(t, "未找到日志", status.Convert(err).Message())
}

type streamCopyToPodServer struct {
	container.Container_StreamCopyToPodServer
	err  error
	idx  int
	recv []*container.StreamCopyToPodRequest
}

func (l *streamCopyToPodServer) Send(response *container.StreamCopyToPodResponse) error {
	return nil
}

func (l *streamCopyToPodServer) SendAndClose(response *container.StreamCopyToPodResponse) error {
	return nil
}

func (l *streamCopyToPodServer) Recv() (*container.StreamCopyToPodRequest, error) {
	if l.err != nil {
		return nil, l.err
	}
	if l.idx < len(l.recv) {
		l.idx++
		return l.recv[l.idx-1], nil
	}
	return nil, io.EOF
}

func (l *streamCopyToPodServer) Context() context.Context {
	return newAdminUserCtx()
}

func TestContainerSvc_StreamCopyToPod_PodNotRunning(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(false, "")
	err := svc.StreamCopyToPod(&streamCopyToPodServer{recv: []*container.StreamCopyToPodRequest{
		{
			Namespace: "a",
			Pod:       "b",
			Container: "c",
		},
	}})
	assert.NotNil(t, err)
}

func TestContainerSvc_StreamCopyToPod_Error(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	eventRepo := repo.NewMockEventRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		eventRepo,
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	err := svc.StreamCopyToPod(&streamCopyToPodServer{err: errors.New("xx")})
	assert.Equal(t, "xx", err.Error())
}

func TestContainerSvc_StreamCopyToPod_Success(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	eventRepo := repo.NewMockEventRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		eventRepo,
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	eventRepo.EXPECT().FileAuditLog(
		types.EventActionType_Upload,
		MustGetUser(newAdminUserCtx()).Name,
		gomock.Any(),
		1,
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(true, "")
	k8sRepo.EXPECT().FindDefaultContainer(gomock.Any(), "a", "b").Return("c", nil)
	fileRepo.EXPECT().StreamUploadFile(gomock.Any(), gomock.Cond(func(x any) bool {
		v := x.(*repo.StreamUploadFileRequest)
		return v.Namespace == "a" &&
			v.Pod == "b" &&
			v.Container == "c" &&
			v.Username == MustGetUser(newAdminUserCtx()).Name &&
			v.FileName == "a.txt"
	})).Return(&repo.File{
		ID:        1,
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	}, nil)
	k8sRepo.EXPECT().CopyFileToPod(gomock.Any(), &repo.CopyFileToPodInput{
		FileId:    1,
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	}).Return(&repo.File{}, nil)
	err := svc.StreamCopyToPod(&streamCopyToPodServer{recv: []*container.StreamCopyToPodRequest{
		{
			Namespace: "a",
			Pod:       "b",
			Container: "",
			FileName:  "a.txt",
			Data:      []byte("data"),
		},
	}})
	assert.Nil(t, err)
}

func TestSortEvents(t *testing.T) {
	event1 := &eventv1.Event{
		ObjectMeta: metav1.ObjectMeta{
			ResourceVersion: "1",
		},
	}
	event2 := &eventv1.Event{
		ObjectMeta: metav1.ObjectMeta{
			ResourceVersion: "2",
		},
	}
	event3 := &eventv1.Event{
		ObjectMeta: metav1.ObjectMeta{
			ResourceVersion: "3",
		},
	}

	events := sortEvents{event3, event1, event2}
	sort.Sort(events)

	assert.Equal(t, "1", events[0].ResourceVersion)
	assert.Equal(t, "2", events[1].ResourceVersion)
	assert.Equal(t, "3", events[2].ResourceVersion)
}

type execOnceServer struct {
	container.Container_ExecOnceServer
	res   []string
	Error *container.ExecError
}

func (l *execOnceServer) Context() context.Context {
	return newAdminUserCtx()
}

func (l *execOnceServer) Send(response *container.ExecResponse) error {
	l.res = append(l.res, string(response.Message))
	l.Error = response.Error
	return nil
}

func TestContainerSvc_ExecOnce_PodNotRunning(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		repo.NewMockFileRepo(m),
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(false, "")
	err := svc.ExecOnce(&container.ExecOnceRequest{
		Namespace: "a",
		Pod:       "b",
	}, &execOnceServer{})
	assert.NotNil(t, err)
}

func TestContainerSvc_ExecOnce_Success(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	eventRepo := repo.NewMockEventRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		eventRepo,
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(true, "")
	k8sRepo.EXPECT().FindDefaultContainer(gomock.Any(), "a", "b").Return("c", nil)
	eventRepo.EXPECT().AuditLogWithChange(
		types.EventActionType_Exec,
		"admin",
		gomock.Any(),
		nil,
		gomock.Cond(func(x any) bool {
			v := x.(repo.AnyYamlPrettier)
			return v["namespace"].(string) == "a" &&
				v["pod"].(string) == "b" &&
				v["container"].(string) == "c" &&
				slices.Equal(v["command"].([]string), []string{"ls"}) &&
				v["error"] == "xx" &&
				v["result"] == ""
		}),
	)

	mac := &execOnceMatcher{
		tty: false,
		cmd: []string{"ls"},
	}
	k8sRepo.EXPECT().Execute(gomock.Any(), &repo.Container{
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	}, mac).Return(clientgoexec.CodeExitError{
		Err:  errors.New("xx"),
		Code: 1,
	})
	ser := &execOnceServer{}
	err := svc.ExecOnce(&container.ExecOnceRequest{
		Namespace: "a",
		Pod:       "b",
		Command:   []string{"ls"},
	}, ser)
	assert.Error(t, err)
	assert.Equal(t, int64(1), ser.Error.Code)
	assert.Equal(t, "xx", ser.Error.Message)
}

type execOnceMatcher struct {
	input *repo.ExecuteInput
	tty   bool
	cmd   []string
}

func (e *execOnceMatcher) Matches(x any) bool {
	input, ok := x.(*repo.ExecuteInput)
	if !ok {
		return false
	}
	e.input = input
	if e.tty != input.TTY {
		return false
	}
	if !slices.Equal(e.cmd, input.Cmd) {
		return false
	}
	return true
}

func (e *execOnceMatcher) String() string {
	return ""
}

func TestContainerSvc_Exec_PodNotRunning(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		repo.NewMockEventRepo(m),
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(false, "Pod not running")
	err := svc.Exec(&execServerMock{})
	assert.NotNil(t, err)
}

func TestContainerSvc_Exec_Success(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	k8sRepo := repo.NewMockK8sRepo(m)
	fileRepo := repo.NewMockFileRepo(m)
	eventRepo := repo.NewMockEventRepo(m)
	svc := NewContainerSvc(
		timer.NewReal(),
		eventRepo,
		k8sRepo,
		fileRepo,
		mlog.NewForConfig(nil),
	)
	eventRepo.EXPECT().FileAuditLogWithDuration(
		types.EventActionType_Exec,
		"mars",
		gomock.Any(),
		1,
		time.Second,
	)
	k8sRepo.EXPECT().IsPodRunning("a", "b").Return(true, "")
	k8sRepo.EXPECT().FindDefaultContainer(gomock.Any(), "a", "b").Return("c", nil)
	reco := &recorderMock{}
	fileRepo.EXPECT().NewRecorder(gomock.Any(), &repo.Container{
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	}).Return(reco)
	k8sRepo.EXPECT().Execute(gomock.Any(), &repo.Container{
		Namespace: "a",
		Pod:       "b",
		Container: "c",
	}, gomock.Cond(func(x any) bool {
		v := x.(*repo.ExecuteInput)
		return v.TTY &&
			slices.Equal(v.Cmd, []string{"ls"}) &&
			v.Stdin != nil &&
			v.Stdout != nil &&
			v.Stderr != nil &&
			v.TerminalSizeQueue != nil
	})).Return(&clientgoexec.CodeExitError{
		Err:  errors.New("xx"),
		Code: 2,
	})
	mock := &execServerMock{}
	err := svc.Exec(mock)
	assert.Equal(t, int64(2), mock.err.Code)
	assert.NotNil(t, err)
	assert.Equal(t, uint16(10), reco.w)
	assert.Equal(t, uint16(20), reco.h)
}

type execServerMock struct {
	container.Container_ExecServer
	err *container.ExecError
}

func (e *execServerMock) Recv() (*container.ExecRequest, error) {
	return &container.ExecRequest{
		Namespace: "a",
		Pod:       "b",
		Command:   []string{"ls"},
		SizeQueue: &container.TerminalSize{
			Width:  10,
			Height: 20,
		},
	}, nil
}

func (e *execServerMock) Send(response *container.ExecResponse) error {
	if response.Error != nil {
		e.err = response.Error
	}
	return nil
}

func (e *execServerMock) Context() context.Context {
	return context.TODO()
}

type recorderMock struct {
	repo.Recorder
	w uint16
	h uint16
}

func (r *recorderMock) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (r *recorderMock) Resize(width, height uint16) {
	r.h = height
	r.w = width
}

func (r *recorderMock) Close() error {
	return nil
}

func (r *recorderMock) Container() *repo.Container {
	return &repo.Container{}
}

func (r *recorderMock) User() *auth.UserInfo {
	return &auth.UserInfo{Name: "mars"}
}

func (r *recorderMock) File() *repo.File {
	return &repo.File{ID: 1}
}

func (r *recorderMock) Duration() time.Duration {
	return time.Second
}

func TestScannerText_SingleLine(t *testing.T) {
	var result string
	err := scannerText("single line", func(s string) {
		result = s
	})
	assert.Nil(t, err)
	assert.Equal(t, "single line", result)
}

func TestScannerText_MultipleLines(t *testing.T) {
	var result []string
	err := scannerText("line1\nline2\nline3", func(s string) {
		result = append(result, s)
	})
	assert.Nil(t, err)
	assert.Equal(t, []string{"line1", "line2", "line3"}, result)
}

func TestScannerText_EmptyString(t *testing.T) {
	var result []string
	err := scannerText("", func(s string) {
		result = append(result, s)
	})
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func TestSizeQueue_Next_ContextDone(t *testing.T) {
	queue := &sizeQueue{
		ch:  make(chan *remotecommand.TerminalSize, 1),
		ctx: context.TODO(),
	}

	ctx, cancel := context.WithCancel(context.TODO())
	queue.ctx = ctx
	cancel()

	assert.Nil(t, queue.Next())
}
func TestSizeQueue_Next_NotOk(t *testing.T) {
	queue := &sizeQueue{
		ch:  make(chan *remotecommand.TerminalSize, 1),
		ctx: context.TODO(),
	}
	close(queue.ch)

	assert.Nil(t, queue.Next())
}

func TestSizeQueue_Next_SizeReceived(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	recorder := repo.NewMockRecorder(m)
	queue := &sizeQueue{
		ch:  make(chan *remotecommand.TerminalSize, 1),
		ctx: context.TODO(),
		r:   recorder,
	}

	expectedSize := &remotecommand.TerminalSize{Width: 10, Height: 20}
	queue.ch <- expectedSize

	recorder.EXPECT().Resize(uint16(expectedSize.Width), uint16(expectedSize.Height))
	assert.Equal(t, expectedSize, queue.Next())
}

func Test_toErrStr(t *testing.T) {
	assert.Equal(t, "", toErrStr(nil))
	assert.Equal(t, "error", toErrStr(errors.New("error")))
}
