package bootstrappers

import (
	"github.com/duc-cnzj/mars/v5/internal/application"
)

type K8sBootstrapper struct{}

func (d *K8sBootstrapper) Tags() []string {
	return []string{}
}

func (d *K8sBootstrapper) Bootstrap(appli application.App) error {
	if appli.Config().IsK8sEnv() {
		return appli.Data().InitK8s(appli.Done())
	}
	return nil
}
