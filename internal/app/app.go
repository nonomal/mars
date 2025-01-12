package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/duc-cnzj/mars/internal/app/bootstrappers"
	app "github.com/duc-cnzj/mars/internal/app/helper"
	"github.com/duc-cnzj/mars/internal/app/instance"
	"github.com/duc-cnzj/mars/internal/cache"
	"github.com/duc-cnzj/mars/internal/config"
	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/database"
	"github.com/duc-cnzj/mars/internal/event"
	"github.com/duc-cnzj/mars/internal/mlog"
	"github.com/duc-cnzj/mars/internal/utils/singleflight"
)

type Hook string

const (
	BeforeRunHook  Hook = "before_run"
	BeforeDownHook Hook = "before_down"
	AfterDownHook  Hook = "after_down"
)

var _ contracts.ApplicationInterface = (*Application)(nil)

var DefaultBootstrappers = []contracts.Bootstrapper{
	&bootstrappers.EventBootstrapper{},
	&bootstrappers.PluginsBootstrapper{},
	&bootstrappers.AuthBootstrapper{},
	&bootstrappers.UploadBootstrapper{},
	&bootstrappers.CacheBootstrapper{},
	&bootstrappers.K8sClientBootstrapper{},
	&bootstrappers.DBBootstrapper{},
	&bootstrappers.ApiGatewayBootstrapper{},
	&bootstrappers.PprofBootstrapper{},
	&bootstrappers.GrpcBootstrapper{},
	&bootstrappers.MetricsBootstrapper{},
	&bootstrappers.OidcBootstrapper{},
	&bootstrappers.TracingBootstrapper{},
	&bootstrappers.AppBootstrapper{},
}

type emptyMetrics struct{}

func (e *emptyMetrics) IncWebsocketConn() {
}

func (e *emptyMetrics) DecWebsocketConn() {
}

type Application struct {
	done          context.Context
	doneFunc      func()
	config        *config.Config
	clientSet     *contracts.K8sClient
	dbManager     contracts.DBManager
	dispatcher    contracts.DispatcherInterface
	metrics       contracts.Metrics
	servers       []contracts.Server
	bootstrappers []contracts.Bootstrapper

	hooksMu sync.RWMutex
	hooks   map[Hook][]contracts.Callback

	plugins      map[string]contracts.PluginInterface
	oidcProvider contracts.OidcConfig
	uploader     contracts.Uploader
	auth         contracts.AuthInterface

	sf    *singleflight.Group
	cache contracts.CacheInterface
}

func (app *Application) SetCache(c contracts.CacheInterface) {
	app.cache = c
}

func (app *Application) Cache() contracts.CacheInterface {
	return app.cache
}

func (app *Application) Auth() contracts.AuthInterface {
	return app.auth
}

func (app *Application) SetAuth(auth contracts.AuthInterface) {
	app.auth = auth
}

func (app *Application) SetUploader(uploader contracts.Uploader) {
	app.uploader = uploader
}

func (app *Application) Uploader() contracts.Uploader {
	return app.uploader
}

func (app *Application) Oidc() contracts.OidcConfig {
	return app.oidcProvider
}

func (app *Application) SetOidc(provider contracts.OidcConfig) {
	app.oidcProvider = provider
}

func (app *Application) SetMetrics(metrics contracts.Metrics) {
	app.metrics = metrics
}

func (app *Application) Metrics() contracts.Metrics {
	return app.metrics
}

func (app *Application) GetPluginByName(name string) contracts.PluginInterface {
	return app.plugins[name]
}

func (app *Application) SetPlugins(plugins map[string]contracts.PluginInterface) {
	app.plugins = plugins
}

func (app *Application) GetPlugins() map[string]contracts.PluginInterface {
	return app.plugins
}

func (app *Application) Done() <-chan struct{} {
	return app.done.Done()
}

func (app *Application) K8sClient() *contracts.K8sClient {
	return app.clientSet
}

func (app *Application) SetK8sClient(client *contracts.K8sClient) {
	app.clientSet = client
}

func (app *Application) EventDispatcher() contracts.DispatcherInterface {
	return app.dispatcher
}

func (app *Application) Singleflight() *singleflight.Group {
	return app.sf
}

func (app *Application) SetEventDispatcher(dispatcher contracts.DispatcherInterface) {
	app.dispatcher = dispatcher
}

type Option func(*Application)

func WithBootstrappers(bootstrappers ...contracts.Bootstrapper) Option {
	return func(app *Application) {
		app.bootstrappers = bootstrappers
	}
}

func NewApplication(config *config.Config, opts ...Option) contracts.ApplicationInterface {
	var mustBooted = []contracts.Bootstrapper{
		&bootstrappers.LogBootstrapper{},
	}

	doneCtx, cancelFunc := context.WithCancel(context.Background())
	app := &Application{
		bootstrappers: DefaultBootstrappers,
		config:        config,
		done:          doneCtx,
		doneFunc:      cancelFunc,
		hooks:         map[Hook][]contracts.Callback{},
		servers:       []contracts.Server{},
		metrics:       &emptyMetrics{},
		sf:            &singleflight.Group{},
		cache:         &cache.NoCache{},
	}

	app.dispatcher = event.NewDispatcher(app)
	app.dbManager = database.NewManager(app)

	for _, opt := range opts {
		opt(app)
	}

	instance.SetInstance(app)

	for _, bootstrapper := range mustBooted {
		if err := bootstrapper.Bootstrap(app); err != nil {
			mlog.Fatal(err)
		}
	}

	if app.IsDebug() {
		printConfig()
	}

	return app
}

func printConfig() {
	mlog.Debugf("imagepullsecrets %#v", app.Config().ImagePullSecrets)
}

func (app *Application) Bootstrap() error {
	for _, bootstrapper := range app.bootstrappers {
		if err := bootstrapper.Bootstrap(app); err != nil {
			return err
		}
	}

	return nil
}

func (app *Application) Config() *config.Config {
	return app.config
}

func (app *Application) DBManager() contracts.DBManager {
	return app.dbManager
}

func (app *Application) IsDebug() bool {
	return app.config.Debug
}

func (app *Application) AddServer(server contracts.Server) {
	app.servers = append(app.servers, server)
}

func (app *Application) Run() context.Context {
	sig := make(chan os.Signal, 2)
	ch, cancel := context.WithCancel(context.TODO())
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	go func() {
		s1 := <-sig
		cancel()
		mlog.Warningf("收到系统信号 %v, 再次执行 ctrl+c 强制退出!", s1.String())
		s2 := <-sig
		mlog.Warningf("收到 %v 信号，执行强制退出!", s2.String())
		os.Exit(1)
	}()

	app.RunServerHooks(BeforeRunHook)

	for _, server := range app.servers {
		if err := server.Run(context.Background()); err != nil {
			mlog.Fatal(err)
		}
	}

	return ch
}

func (app *Application) Shutdown() {
	app.doneFunc()
	app.RunServerHooks(BeforeDownHook)

	wg := &sync.WaitGroup{}
	for _, server := range app.servers {
		wg.Add(1)
		go func(server contracts.Server) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				mlog.Error(err)
			}
		}(server)
	}
	wg.Wait()

	app.RunServerHooks(AfterDownHook)

	mlog.Info("server graceful shutdown.")
}

func (app *Application) RegisterAfterShutdownFunc(fn contracts.Callback) {
	app.hooksMu.Lock()
	defer app.hooksMu.Unlock()
	app.hooks[AfterDownHook] = append(app.hooks[AfterDownHook], fn)
}

func (app *Application) RegisterBeforeShutdownFunc(fn contracts.Callback) {
	app.hooksMu.Lock()
	defer app.hooksMu.Unlock()
	app.hooks[BeforeDownHook] = append(app.hooks[BeforeDownHook], fn)
}

func (app *Application) RunServerHooks(hook Hook) {
	app.hooksMu.RLock()
	defer app.hooksMu.RUnlock()
	wg := sync.WaitGroup{}
	for _, cb := range app.hooks[hook] {
		wg.Add(1)
		go func(cb contracts.Callback) {
			defer wg.Done()
			cb(app)
		}(cb)
	}
	wg.Wait()
}

func (app *Application) BeforeServerRunHooks(cb contracts.Callback) {
	app.hooksMu.Lock()
	defer app.hooksMu.Unlock()
	app.hooks[BeforeRunHook] = append(app.hooks[BeforeRunHook], cb)
}
