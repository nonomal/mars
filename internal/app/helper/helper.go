package app

import (
	"github.com/duc-cnzj/mars/internal/app/instance"
	"github.com/duc-cnzj/mars/internal/config"
	"github.com/duc-cnzj/mars/internal/contracts"
	"github.com/duc-cnzj/mars/internal/utils/singleflight"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

func App() contracts.ApplicationInterface {
	return instance.App()
}

func Auth() contracts.AuthInterface {
	return App().Auth()
}

func Oidc() contracts.OidcConfig {
	return App().Oidc()
}

func Config() *config.Config {
	return App().Config()
}

func DB() *gorm.DB {
	return App().DBManager().DB()
}

func Uploader() contracts.Uploader {
	return App().Uploader()
}

func Event() contracts.DispatcherInterface {
	return App().EventDispatcher()
}

func K8sClient() *contracts.K8sClient {
	return App().K8sClient()
}

func K8sClientSet() kubernetes.Interface {
	return App().K8sClient().Client
}

func K8sMetrics() versioned.Interface {
	return App().K8sClient().MetricsClient
}

func Metrics() contracts.Metrics {
	return App().Metrics()
}

func Singleflight() *singleflight.Group {
	return App().Singleflight()
}

func Cache() contracts.CacheInterface {
	return App().Cache()
}
