package application

import (
	"context"
	"net/http"

	"github.com/duc-cnzj/mars/v5/internal/auth"
	"github.com/duc-cnzj/mars/v5/internal/cache"
	"github.com/duc-cnzj/mars/v5/internal/config"
	"github.com/duc-cnzj/mars/v5/internal/cron"
	"github.com/duc-cnzj/mars/v5/internal/data"
	"github.com/duc-cnzj/mars/v5/internal/ent"
	"github.com/duc-cnzj/mars/v5/internal/event"
	"github.com/duc-cnzj/mars/v5/internal/locker"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/uploader"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
)

type EndpointFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
type RegistryFunc = func(s grpc.ServiceRegistrar)

type GrpcRegistry struct {
	EndpointFuncs []EndpointFunc
	RegistryFunc  RegistryFunc
}

type Callback func(App)

// Server define booting server.
type Server interface {
	// Run server.
	Run(context.Context) error
	// Shutdown server.
	Shutdown(context.Context) error
}

// Bootstrapper boots.
type Bootstrapper interface {
	// Bootstrap when app start.
	Bootstrap(App) error
	// Tags boot tags.
	Tags() []string
}

// App app.
type App interface {
	// Data app data.
	Data() data.Data

	// Config app configuration.
	Config() *config.Config

	// IsDebug bool.
	IsDebug() bool

	// GrpcRegistry return register.
	GrpcRegistry() *GrpcRegistry

	// Logger return logger.
	Logger() mlog.Logger

	// Auth return repo.Auth.
	Auth() auth.Auth

	// Oidc return oidc config.
	Oidc() data.OidcConfig

	// PrometheusRegistry return prometheus.
	PrometheusRegistry() *prometheus.Registry

	// Uploader getter
	Uploader() uploader.Uploader

	// Bootstrap boots all.
	Bootstrap() error

	// DB instance.
	DB() *ent.Client

	// AddServer add boot server
	AddServer(Server)

	// Run servers.
	Run() context.Context

	// Shutdown all servers.
	Shutdown()

	// Done return done chan.
	Done() <-chan struct{}

	// BeforeServerRunHooks register hooks.
	BeforeServerRunHooks(Callback)

	// RegisterBeforeShutdownFunc register hooks.
	RegisterBeforeShutdownFunc(Callback)

	// RegisterAfterShutdownFunc register hooks.
	RegisterAfterShutdownFunc(Callback)

	// Dispatcher return eventer.
	Dispatcher() event.Dispatcher

	// PluginMgr return plugin manager.
	PluginMgr() PluginManger

	// Singleflight return singleflight.
	Singleflight() *singleflight.Group

	// Cache return cache.
	Cache() cache.Cache

	// Locker return locker
	Locker() locker.Locker

	// CronManager return cron manager
	CronManager() cron.Manager

	HttpHandler() HttpHandler
}

type WsHttpServer interface {
	TickClusterHealth(done <-chan struct{})
	Info(writer http.ResponseWriter, request *http.Request)
	Serve(w http.ResponseWriter, r *http.Request)
	Shutdown(ctx context.Context) error
}

type HttpHandler interface {
	WsHttpServer

	RegisterWsRoute(mux *mux.Router)
	RegisterSwaggerUIRoute(mux *mux.Router)
	RegisterFileRoute(mux *runtime.ServeMux)
}
