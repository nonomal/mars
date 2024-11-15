package api

import (
	"context"
	"crypto/tls"
	"io"
	"strings"
	"sync/atomic"

	"github.com/cenkalti/backoff/v4"
	"github.com/duc-cnzj/mars/api/v5/auth"
	"github.com/duc-cnzj/mars/api/v5/changelog"
	"github.com/duc-cnzj/mars/api/v5/cluster"
	"github.com/duc-cnzj/mars/api/v5/container"
	"github.com/duc-cnzj/mars/api/v5/endpoint"
	"github.com/duc-cnzj/mars/api/v5/event"
	"github.com/duc-cnzj/mars/api/v5/file"
	"github.com/duc-cnzj/mars/api/v5/git"
	"github.com/duc-cnzj/mars/api/v5/metrics"
	"github.com/duc-cnzj/mars/api/v5/namespace"
	"github.com/duc-cnzj/mars/api/v5/picture"
	"github.com/duc-cnzj/mars/api/v5/project"
	"github.com/duc-cnzj/mars/api/v5/repo"
	"github.com/duc-cnzj/mars/api/v5/token"
	"github.com/duc-cnzj/mars/api/v5/version"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interface interface {
	io.Closer
	SetBearerToken(string)

	Auth() auth.AuthClient
	Picture() picture.PictureClient
	Version() version.VersionClient
	Cluster() cluster.ClusterClient
	Changelog() changelog.ChangelogClient
	Event() event.EventClient
	Container() container.ContainerClient
	File() file.FileClient
	Git() git.GitClient
	Namespace() namespace.NamespaceClient
	Project() project.ProjectClient
	Endpoint() endpoint.EndpointClient
	Metrics() metrics.MetricsClient
	AccessToken() token.AccessTokenClient
	Repo() repo.RepoClient
}

type Client struct {
	singleflight Group

	UnaryClientInterceptors  []grpc.UnaryClientInterceptor
	StreamClientInterceptors []grpc.StreamClientInterceptor
	username, password       string
	authTokenValue           atomic.Value

	tls *tls.Config

	conn        *grpc.ClientConn
	dialOptions []grpc.DialOption

	auth        auth.AuthClient
	changelog   changelog.ChangelogClient
	cluster     cluster.ClusterClient
	container   container.ContainerClient
	endpoint    endpoint.EndpointClient
	event       event.EventClient
	file        file.FileClient
	git         git.GitClient
	metrics     metrics.MetricsClient
	namespace   namespace.NamespaceClient
	picture     picture.PictureClient
	project     project.ProjectClient
	repo        repo.RepoClient
	accessToken token.AccessTokenClient
	version     version.VersionClient
}

var _ Interface = (*Client)(nil)

func NewClient(addr string, opts ...Option) (Interface, error) {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	dial, err := grpc.NewClient(addr, c.buildDialOptions()...)

	if err != nil {
		return nil, err
	}
	c.conn = dial

	c.auth = auth.NewAuthClient(dial)
	c.changelog = changelog.NewChangelogClient(dial)
	c.cluster = cluster.NewClusterClient(dial)
	c.container = container.NewContainerClient(dial)
	c.event = event.NewEventClient(dial)
	c.git = git.NewGitClient(dial)
	c.metrics = metrics.NewMetricsClient(dial)
	c.namespace = namespace.NewNamespaceClient(dial)
	c.picture = picture.NewPictureClient(dial)
	c.project = project.NewProjectClient(dial)
	c.version = version.NewVersionClient(dial)
	c.file = file.NewFileClient(dial)
	c.endpoint = endpoint.NewEndpointClient(dial)
	c.accessToken = token.NewAccessTokenClient(dial)
	c.repo = repo.NewRepoClient(dial)

	if c.password != "" || c.username != "" {
		if err := c.getToken(); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) SetBearerToken(token string) {
	c.setToken(token)
}

func (c *Client) hasCredentials() bool {
	return c.username != "" && c.password != ""
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}

func (c *Client) Auth() auth.AuthClient {
	return c.auth
}

func (c *Client) Repo() repo.RepoClient {
	return c.repo
}

func (c *Client) Changelog() changelog.ChangelogClient {
	return c.changelog
}

func (c *Client) Cluster() cluster.ClusterClient {
	return c.cluster
}

func (c *Client) Container() container.ContainerClient {
	return c.container
}

func (c *Client) Event() event.EventClient {
	return c.event
}

func (c *Client) AccessToken() token.AccessTokenClient {
	return c.accessToken
}

func (c *Client) File() file.FileClient {
	return c.file
}

func (c *Client) Git() git.GitClient {
	return c.git
}

func (c *Client) Metrics() metrics.MetricsClient {
	return c.metrics
}

func (c *Client) Namespace() namespace.NamespaceClient {
	return c.namespace
}

func (c *Client) Picture() picture.PictureClient {
	return c.picture
}

func (c *Client) Project() project.ProjectClient {
	return c.project
}

func (c *Client) Version() version.VersionClient {
	return c.version
}

func (c *Client) Endpoint() endpoint.EndpointClient {
	return c.endpoint
}

func (c *Client) authToken() string {
	v := c.authTokenValue.Load()
	if v != nil {
		return v.(string)
	}
	return ""
}

func (c *Client) buildDialOptions() []grpc.DialOption {
	if c.tls == nil {
		c.dialOptions = append(c.dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if !c.hasCredentials() && c.authTokenValue.Load() == nil {
		c.dialOptions = append(c.dialOptions, grpc.WithPerRPCCredentials(&clientauth{c: c}))
	}

	c.dialOptions = append(c.dialOptions,
		grpc.WithChainStreamInterceptor(c.StreamClientInterceptors...),
		grpc.WithChainUnaryInterceptor(c.UnaryClientInterceptors...),
	)

	return c.dialOptions
}

func (c *Client) getToken() error {
	login, err, _ := c.singleflight.Do("Retry", func() (interface{}, error) {
		return c.auth.Login(context.TODO(), &auth.LoginRequest{
			Username: c.username,
			Password: c.password,
		})
	})
	if err != nil {
		return err
	}

	c.setToken(login.(*auth.LoginResponse).Token)
	return nil
}

func (c *Client) setToken(token string) {
	if len(token) > 6 && !strings.EqualFold("Bearer", token[0:6]) {
		token = "Bearer " + token
	}
	c.authTokenValue.Store(token)
}

type Option func(*Client)

type clientauth struct {
	c *Client
}

func (a *clientauth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": a.c.authToken(),
	}, nil
}

func (a *clientauth) RequireTransportSecurity() bool {
	return false
}

func WithAuth(username, password string) Option {
	return func(c *Client) {
		c.username = username
		c.password = password
		c.dialOptions = append(c.dialOptions, grpc.WithPerRPCCredentials(&clientauth{c: c}))
	}
}

func WithBearerToken(token string) Option {
	return func(c *Client) {
		c.setToken(token)
		c.dialOptions = append(c.dialOptions, grpc.WithPerRPCCredentials(&clientauth{c: c}))
	}
}

// WithTokenAutoRefresh
// TODO c.StreamClientInterceptors 有点难搞，好在目前没用到，之后用到了需要搞一下
func WithTokenAutoRefresh() Option {
	return func(c *Client) {
		c.UnaryClientInterceptors = append(c.UnaryClientInterceptors,
			func(
				ctx context.Context,
				method string,
				req, reply interface{},
				cc *grpc.ClientConn,
				invoker grpc.UnaryInvoker,
				opts ...grpc.CallOption) error {
				err := invoker(ctx, method, req, reply, cc, opts...)
				operation := func() error {
					if gerr := c.getToken(); gerr != nil {
						return gerr
					}
					return invoker(ctx, method, req, reply, cc, opts...)
				}
				var bf backoff.BackOff = backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
				bf = backoff.WithMaxRetries(bf, 5)

				if c.hasCredentials() && status.Code(err) == codes.Unauthenticated && method != "/Auth/Login" {
					return backoff.Retry(operation, bf)
				}
				return err
			})
	}
}

// WithTransportCredentials 暂时不支持
//func WithTransportCredentials(tls *tls.Config) Option {
//	return func(c *Client) {
//		c.tls = tls
//		c.dialOptions = append(c.dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(tls)))
//	}
//}

func WithUnaryClientInterceptor(op grpc.UnaryClientInterceptor) Option {
	return func(c *Client) {
		c.UnaryClientInterceptors = append(c.UnaryClientInterceptors, op)
	}
}

func WithStreamClientInterceptor(op grpc.StreamClientInterceptor) Option {
	return func(c *Client) {
		c.StreamClientInterceptors = append(c.StreamClientInterceptors, op)
	}
}

func WithTracer() Option {
	return func(c *Client) {
		c.dialOptions = append(c.dialOptions, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	}
}

type GatewayCarrier metadata.MD

func (hc GatewayCarrier) Get(key string) string {
	vals := metadata.MD(hc).Get(key)
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func (hc GatewayCarrier) Set(key string, value string) {
	metadata.MD(hc).Set(key, value)
}

func (hc GatewayCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range hc {
		keys = append(keys, k)
	}
	return keys
}
