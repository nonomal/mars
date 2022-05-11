package bootstrappers

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/duc-cnzj/mars/internal/auth"
	"github.com/duc-cnzj/mars/internal/cache"
	"github.com/duc-cnzj/mars/internal/config"
	"github.com/duc-cnzj/mars/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthBootstrapper_Bootstrap(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	app := mock.NewMockApplicationInterface(controller)
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	privateKey, _ := x509.MarshalPKCS8PrivateKey(key)
	bf := bytes.Buffer{}
	pem.Encode(&bf, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKey})
	app.EXPECT().Config().Return(&config.Config{
		PrivateKey: bf.String(),
	}).Times(1)
	app.EXPECT().SetAuth(auth.NewAuth(key, key.Public().(*rsa.PublicKey))).Times(1)
	(&AuthBootstrapper{}).Bootstrap(app)
}

type cacheMatcher struct {
	wants any
	t     *testing.T
}

func (c *cacheMatcher) Matches(x any) bool {
	assert.IsType(c.t, c.wants, x)
	return true
}

func (c *cacheMatcher) String() string {
	return ""
}

func TestCacheBootstrapper_Bootstrap(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	app := mock.NewMockApplicationInterface(controller)
	app.EXPECT().Config().Return(&config.Config{
		DBDriver:    "sqlite",
		CacheDriver: "db",
	}).Times(1)
	app.EXPECT().Singleflight().Times(1)
	app.EXPECT().SetCache(&cacheMatcher{
		wants: (*cache.Cache)(nil),
		t:     t,
	})
	assert.Nil(t, (&CacheBootstrapper{}).Bootstrap(app))
	app.EXPECT().Config().Return(&config.Config{
		DBDriver:    "mysql",
		CacheDriver: "db",
	}).Times(1)
	app.EXPECT().SetCache(&cacheMatcher{
		wants: (*cache.DBCache)(nil),
		t:     t,
	})
	assert.Nil(t, (&CacheBootstrapper{}).Bootstrap(app))
	app.EXPECT().Config().Return(&config.Config{
		CacheDriver: "xxxx",
	}).Times(1)
	assert.Error(t, (&CacheBootstrapper{}).Bootstrap(app))
}

func TestMetricsBootstrapper_Bootstrap(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	app := mock.NewMockApplicationInterface(controller)
	app.EXPECT().SetMetrics(gomock.Any()).Times(1)
	app.EXPECT().AddServer(gomock.Any()).Times(1)
	(&MetricsBootstrapper{}).Bootstrap(app)
}

func TestOidcBootstrapper_Bootstrap(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	app := mock.NewMockApplicationInterface(controller)
	app.EXPECT().Config().Times(1).Return(&config.Config{})
	app.EXPECT().SetOidc(gomock.Any()).Times(1)
	(&OidcBootstrapper{}).Bootstrap(app)
}

func TestPluginsBootstrapper_Bootstrap(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	app := mock.NewMockApplicationInterface(controller)
	app.EXPECT().SetPlugins(gomock.Any()).Times(1)
	(&PluginsBootstrapper{}).Bootstrap(app)
}
func TestGrpcBootstrapper_Bootstrap(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	app := mock.NewMockApplicationInterface(controller)
	app.EXPECT().AddServer(gomock.Any()).Times(1)
	app.EXPECT().Config().Times(1).Return(&config.Config{})
	(&GrpcBootstrapper{}).Bootstrap(app)
}

func TestPprofBootstrapper_Bootstrap(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	app := mock.NewMockApplicationInterface(controller)
	app.EXPECT().AddServer(gomock.Any()).Times(0)
	app.EXPECT().Config().Times(1).Return(&config.Config{
		ProfileEnabled: false,
	})
	(&PprofBootstrapper{}).Bootstrap(app)
	app.EXPECT().AddServer(gomock.Any()).Times(1)
	app.EXPECT().Config().Times(1).Return(&config.Config{
		ProfileEnabled: true,
	})
	(&PprofBootstrapper{}).Bootstrap(app)
}