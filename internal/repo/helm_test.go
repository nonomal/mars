package repo

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/duc-cnzj/mars/api/v5/types"
	websocket_pb "github.com/duc-cnzj/mars/api/v5/websocket"
	"github.com/duc-cnzj/mars/v5/internal/config"
	"github.com/duc-cnzj/mars/v5/internal/data"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
	corev1 "k8s.io/api/core/v1"
	eventv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	restclient "k8s.io/client-go/rest"
)

func TestReleaseList_Add(t *testing.T) {
	rl := ReleaseList{}
	rl.Add(&release.Release{Name: "rl1", Namespace: "dev", Info: &release.Info{Status: "deployed"}})
	rl.Add(&release.Release{Name: "rl2", Namespace: "dev", Info: &release.Info{Status: "pending-upgrade"}})
	rl.Add(&release.Release{Name: "rl3", Namespace: "dev", Info: &release.Info{Status: "pending-rollback"}})
	rl.Add(&release.Release{Name: "rl4", Namespace: "dev", Info: &release.Info{Status: "pending-install"}})
	rl.Add(&release.Release{Name: "rl5", Namespace: "dev", Info: &release.Info{Status: "uninstalling"}})
	rl.Add(&release.Release{Name: "rl6", Namespace: "dev", Info: &release.Info{Status: "failed"}})
	rl.Add(&release.Release{Name: "rl7", Namespace: "dev", Info: &release.Info{Status: "superseded"}})
	rl.Add(&release.Release{Name: "rl8", Namespace: "dev", Info: &release.Info{Status: "unknown"}})
	assert.Len(t, rl, 8)
	_, ok := rl["dev-rl1"]
	assert.True(t, ok)
	assert.Equal(t, "deployed", rl["dev-rl1"].Status)
	assert.Equal(t, "pending", rl["dev-rl2"].Status)
	assert.Equal(t, "pending", rl["dev-rl3"].Status)
	assert.Equal(t, "pending", rl["dev-rl4"].Status)
	assert.Equal(t, "unknown", rl["dev-rl5"].Status)
	assert.Equal(t, "failed", rl["dev-rl6"].Status)
	assert.Equal(t, "unknown", rl["dev-rl7"].Status)
	assert.Equal(t, "unknown", rl["dev-rl8"].Status)
}
func TestReleaseList_GetStatus(t *testing.T) {
	rl := ReleaseList{}
	rl.Add(&release.Release{Name: "rl1", Namespace: "dev", Info: &release.Info{Status: "deployed"}})
	assert.Equal(t, "deployed", rl.GetStatus("dev", "rl1"))
	assert.Equal(t, "unknown", rl.GetStatus("dev", "xxx"))
}

func TestReleaseStatus(t *testing.T) {
	status := (&DefaultHelmer{
		logger: mlog.NewForConfig(nil),
	}).releaseStatus("a", "test", "xx")
	assert.Equal(t, types.Deploy_StatusUnknown, status)
}

func TestRollback(t *testing.T) {
	err := (&DefaultHelmer{}).Rollback("test", "ns", false, nil, false)
	assert.Error(t, err)
}

func TestUninstallRelease(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	err := (&DefaultHelmer{}).Uninstall("test", "ns", func(format string, v ...any) {})
	assert.Error(t, err)
}

func Test_checkIfInstallable(t *testing.T) {
	err := checkIfInstallable(&chart.Chart{
		Metadata: &chart.Metadata{
			Type: "",
		},
	})
	assert.Nil(t, err)
	err = checkIfInstallable(&chart.Chart{
		Metadata: &chart.Metadata{
			Type: "application",
		},
	})
	assert.Nil(t, err)
	err = checkIfInstallable(&chart.Chart{
		Metadata: &chart.Metadata{
			Type: "xxx",
		},
	})
	assert.Error(t, err)
}

func Test_getActionConfigAndSettings(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	settings := getActionConfigAndSettings("test", "", func(format string, v ...any) {})
	assert.NotNil(t, settings)
}

func Test_watchEvent(t *testing.T) {
	t.Parallel()
	ctx, cancelFn := context.WithCancel(context.TODO())
	ch := make(chan data.Obj[*eventv1.Event], 10)
	go func() {
		ch <- data.NewObj(nil, &eventv1.Event{
			Regarding: corev1.ObjectReference{
				Namespace: "ns",
				Name:      "app",
			},
		}, data.Add)
		ch <- data.NewObj(nil, &eventv1.Event{
			Regarding: corev1.ObjectReference{
				Namespace: "ns",
				Name:      "app1",
			},
		}, data.Update)
		ch <- data.NewObj(nil, &eventv1.Event{
			Regarding: corev1.ObjectReference{
				Namespace: "ns",
				Name:      "app2",
			},
		}, data.Delete)
		time.Sleep(2 * time.Second)
		cancelFn()
	}()
	var called int64
	(&DefaultHelmer{
		logger: mlog.NewForConfig(nil),
	}).watchEvent(ctx, ch, "release", func(container []*websocket_pb.Container, format string, v ...any) {
		atomic.AddInt64(&called, 1)
	}, NewPodLister(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns",
			Name:      "app",
			Labels: map[string]string{
				"xxx":          "xxx",
				"release-name": "release",
			},
		},
	}))

	assert.Equal(t, int64(1), atomic.LoadInt64(&called))
}

func Test_watchEvent_Error1(t *testing.T) {
	t.Parallel()
	ctx, cancelFn := context.WithCancel(context.TODO())
	ch := make(chan data.Obj[*eventv1.Event], 10)
	go func() {
		close(ch)
		time.Sleep(2 * time.Second)
		cancelFn()
	}()
	var called int64
	(&DefaultHelmer{
		logger: mlog.NewForConfig(nil),
	}).watchEvent(ctx, ch, "release", func(container []*websocket_pb.Container, format string, v ...any) {
		atomic.AddInt64(&called, 1)
	}, NewPodLister(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns",
			Name:      "app",
			Labels: map[string]string{
				"xxx":          "xxx",
				"release-name": "release",
			},
		},
	}))

	assert.Equal(t, int64(0), atomic.LoadInt64(&called))
}

func Test_watchEvent_Error2(t *testing.T) {
	t.Parallel()
	ctx, cancelFn := context.WithCancel(context.TODO())
	ch := make(chan data.Obj[*eventv1.Event], 10)
	go func() {
		ch <- data.NewObj(nil, &eventv1.Event{
			Regarding: corev1.ObjectReference{
				Namespace: "ns",
				Name:      "app1",
			},
		}, data.Add)
		ch <- data.NewObj(nil, &eventv1.Event{
			Regarding: corev1.ObjectReference{
				Namespace: "ns",
				Name:      "app",
			},
		}, data.Add)
		time.Sleep(2 * time.Second)
		cancelFn()
	}()
	var called int64
	(&DefaultHelmer{
		logger: mlog.NewForConfig(nil),
	}).watchEvent(ctx, ch, "release", func(container []*websocket_pb.Container, format string, v ...any) {
		atomic.AddInt64(&called, 1)
	}, NewPodLister(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns",
			Name:      "app",
			Labels: map[string]string{
				"xxx":          "xxx",
				"release-name": "release",
			},
		},
	}))

	assert.Equal(t, int64(1), atomic.LoadInt64(&called))
}

func Test_watchPodStatus(t *testing.T) {
	t.Parallel()
	var called int64
	podCh := make(chan data.Obj[*corev1.Pod], 10)
	ctx, cancelFn := context.WithCancel(context.TODO())
	go func() {
		podCh <- data.NewObj(nil, &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "ns",
				Name:      "app1",
				Labels: map[string]string{
					"name": "app",
				},
			},
			Status: corev1.PodStatus{
				ContainerStatuses: []corev1.ContainerStatus{
					{
						Name:         "aaa",
						Ready:        false,
						RestartCount: 6,
					},
				},
			},
		}, data.Add)
		podCh <- data.NewObj(nil, &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "ns",
				Name:      "app2",
				Labels: map[string]string{
					"name": "app",
				},
			},
			Status: corev1.PodStatus{
				ContainerStatuses: []corev1.ContainerStatus{
					{
						Name:         "bbb",
						Ready:        false,
						RestartCount: 5,
					},
				},
			},
		}, data.Delete)
		podCh <- data.NewObj[*corev1.Pod](nil, nil, data.Update)
		podCh <- data.NewObj[*corev1.Pod](nil, nil, data.Update)
		podCh <- data.NewObj[*corev1.Pod](nil, nil, data.Update)
		time.Sleep(2 * time.Second)
		cancelFn()
	}()
	selectorLists := []labels.Selector{
		labels.SelectorFromSet(map[string]string{
			"name": "app",
		}),
		labels.SelectorFromSet(map[string]string{
			"release": "v1",
		}),
	}
	(&DefaultHelmer{
		logger: mlog.NewForConfig(nil),
	}).watchPodStatus(ctx, podCh, selectorLists, func(container []*websocket_pb.Container, format string, v ...any) {
		atomic.AddInt64(&called, 1)
	})
	assert.Equal(t, int64(2), atomic.LoadInt64(&called))

	podCh2 := make(chan data.Obj[*corev1.Pod], 10)
	close(podCh2)
	assert.NotPanics(t, func() {
		(&DefaultHelmer{
			logger: mlog.NewForConfig(nil),
		}).watchPodStatus(context.TODO(), podCh2, nil, nil)
	})
}

func Test_watchPodStatus_Error1(t *testing.T) {
	t.Parallel()
	var called int64
	var cs = &ContainerGetterSetter{}
	podCh := make(chan data.Obj[*corev1.Pod], 10)
	ctx, cancelFn := context.WithCancel(context.TODO())
	go func() {
		podCh <- data.NewObj[*corev1.Pod](nil, &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "ns",
				Name:      "app-not-match",
				Labels: map[string]string{
					"name": "app-not-match",
				},
			},
		}, data.Add)
		podCh <- data.NewObj[*corev1.Pod](nil, &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "ns",
				Name:      "app",
				Labels: map[string]string{
					"name": "app",
				},
			},
			Status: corev1.PodStatus{
				ContainerStatuses: []corev1.ContainerStatus{
					{
						Name:         "three",
						Ready:        false,
						RestartCount: 0,
					},
				},
			},
		}, data.Delete)
		time.Sleep(2 * time.Second)
		cancelFn()
	}()
	selectorLists := []labels.Selector{
		labels.SelectorFromSet(map[string]string{
			"name": "app",
		}),
		labels.SelectorFromSet(map[string]string{
			"release": "v1",
		}),
	}
	(&DefaultHelmer{
		logger: mlog.NewForConfig(nil),
	}).watchPodStatus(ctx, podCh, selectorLists, func(container []*websocket_pb.Container, format string, v ...any) {
		cs.Set(container)
		atomic.AddInt64(&called, 1)
	})
	assert.Len(t, cs.Get(), 0)
	assert.Equal(t, int64(0), atomic.LoadInt64(&called))
}

type ContainerGetterSetter struct {
	sync.Mutex
	cs []*websocket_pb.Container
}

func (c *ContainerGetterSetter) Set(cs []*websocket_pb.Container) {
	c.Lock()
	defer c.Unlock()
	c.cs = cs
}
func (c *ContainerGetterSetter) Get() []*websocket_pb.Container {
	c.Lock()
	defer c.Unlock()
	return c.cs
}

func Test_formatStatus(t *testing.T) {
	var tests = []struct {
		input release.Status
		want  types.Deploy
	}{
		{
			input: release.StatusPendingUpgrade,
			want:  types.Deploy_StatusDeploying,
		},
		{
			input: release.StatusPendingInstall,
			want:  types.Deploy_StatusDeploying,
		},
		{
			input: release.StatusPendingRollback,
			want:  types.Deploy_StatusDeploying,
		},
		{
			input: release.StatusDeployed,
			want:  types.Deploy_StatusDeployed,
		},
		{
			input: release.StatusFailed,
			want:  types.Deploy_StatusFailed,
		},
		{
			input: "xxx",
			want:  types.Deploy_StatusUnknown,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run("", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, formatStatus(tt.input))
		})
	}
}

func Test_fillInstall(t *testing.T) {
	i := &action.Install{}
	u := &action.Upgrade{
		Install:                  true,
		Devel:                    true,
		Namespace:                "xxx",
		SkipCRDs:                 true,
		Timeout:                  10,
		Wait:                     true,
		WaitForJobs:              true,
		DisableHooks:             true,
		DryRun:                   true,
		Force:                    true,
		Atomic:                   true,
		SubNotes:                 true,
		Description:              "desc",
		DisableOpenAPIValidation: true,
		DependencyUpdate:         true,
	}
	fillInstall(i, u)

	assert.Equal(t, i.CreateNamespace, true)
	assert.Equal(t, i.ChartPathOptions, u.ChartPathOptions)
	assert.Equal(t, i.DryRun, u.DryRun)
	assert.Equal(t, i.DisableHooks, u.DisableHooks)
	assert.Equal(t, i.SkipCRDs, u.SkipCRDs)
	assert.Equal(t, i.Timeout, u.Timeout)
	assert.Equal(t, i.Wait, u.Wait)
	assert.Equal(t, i.WaitForJobs, u.WaitForJobs)
	assert.Equal(t, i.Devel, u.Devel)
	assert.Equal(t, i.Namespace, u.Namespace)
	assert.Equal(t, i.Atomic, u.Atomic)
	assert.Equal(t, i.PostRenderer, u.PostRenderer)
	assert.Equal(t, i.DisableOpenAPIValidation, u.DisableOpenAPIValidation)
	assert.Equal(t, i.SubNotes, u.SubNotes)
	assert.Equal(t, i.Description, u.Description)
	assert.Equal(t, i.DependencyUpdate, u.DependencyUpdate)
}

func Test_wrapRestConfig(t *testing.T) {
	cfg := &restclient.Config{}
	wrapRestConfig(cfg)
	assert.Equal(t, float32(-1), cfg.QPS)
}

func Test_logWriter_Write(t *testing.T) {
	n, err := (&logWriter{}).Write([]byte("ass"))
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
}

func Test_newDefaultRegistryClient(t *testing.T) {
	client, err := newDefaultRegistryClient(false, "")
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestWrapLogFn_UnWrap(t *testing.T) {
	called := false
	WrapLogFn(func(container []*websocket_pb.Container, format string, v ...any) {
		called = true
	})(nil, "", "")
	assert.True(t, called)
}

func TestNewDefaultHelmer(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	mockData := data.NewMockData(m)
	k8sRepo := NewMockK8sRepo(m)
	helmer := NewDefaultHelmer(k8sRepo, mockData, &config.Config{}, mlog.NewForConfig(nil)).(*DefaultHelmer)
	assert.NotNil(t, helmer.data)
	assert.NotNil(t, helmer.logger)
	assert.NotNil(t, helmer.k8sRepo)
}
