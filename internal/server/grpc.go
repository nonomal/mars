package server

import (
	"context"
	"net"

	"github.com/duc-cnzj/mars/v5/internal/application"
	marsauthorizor "github.com/duc-cnzj/mars/v5/internal/auth"
	"github.com/duc-cnzj/mars/v5/internal/mlog"
	"github.com/duc-cnzj/mars/v5/internal/server/middlewares"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServerImp interface {
	GracefulStop()
	Serve(lis net.Listener) error
}

type grpcRunner struct {
	server       GrpcServerImp
	endpoint     string
	logger       mlog.Logger
	auth         marsauthorizor.Auth
	grpcRegistry *application.GrpcRegistry
}

func NewGrpcRunner(
	endpoint string,
	app application.App,
) application.Server {
	return &grpcRunner{
		grpcRegistry: app.GrpcRegistry(),
		endpoint:     endpoint,
		logger:       app.Logger().WithModule("server/grpcRunner"),
		auth:         app.Auth(),
	}
}

func (g *grpcRunner) Shutdown(ctx context.Context) error {
	defer g.logger.Info("[Server]: shutdown grpcRunner runner.")
	done := make(chan struct{})
	go func() {
		g.server.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (g *grpcRunner) Run(ctx context.Context) error {
	g.logger.Infof("[Server]: start grpcRunner runner at %s.", g.endpoint)
	listen, err := net.Listen("tcp", g.endpoint)
	if err != nil {
		return err
	}
	g.server = g.initServer()
	go func() {
		if err := g.server.Serve(listen); err != nil {
			g.logger.Error(err)
		}
	}()

	return nil
}

func (g *grpcRunner) initServer() *grpc.Server {
	authFn := func(ctx context.Context) (context.Context, error) {
		return authenticate(ctx, g.auth)
	}
	server := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainStreamInterceptor(
			grpc_auth.StreamServerInterceptor(authFn),
			middlewares.AuthStreamServerInterceptor(),
			middlewares.ValidatorStreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(g.recoveryHandler)),
			middlewares.MetricsStreamServerInterceptor(g.logger),
		),
		grpc.ChainUnaryInterceptor(
			middlewares.LoggerUnaryServerInterceptor(g.logger),
			grpc_auth.UnaryServerInterceptor(authFn),
			middlewares.MetricsServerInterceptor(g.logger),
			middlewares.AuthUnaryServerInterceptor(),
			middlewares.ValidatorUnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(g.recoveryHandler)),
		),
	)

	g.grpcRegistry.RegistryFunc(server)

	return server
}

func (g *grpcRunner) recoveryHandler(p any) error {
	g.logger.Errorf("[Grpc]: recovery error: \n%v", p)
	return nil
}

func authenticate(ctx context.Context, auth marsauthorizor.Auth) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	if verifyToken, b := auth.VerifyToken(token); b {
		return marsauthorizor.SetUser(ctx, verifyToken.UserInfo), nil
	}

	return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated.")
}
