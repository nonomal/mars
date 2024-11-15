package bootstrappers

import (
	"github.com/duc-cnzj/mars/v5/internal/application"
	"github.com/duc-cnzj/mars/v5/internal/server"
)

type MetricsBootstrapper struct{}

func (m *MetricsBootstrapper) Tags() []string {
	return []string{"metrics"}
}

func (m *MetricsBootstrapper) Bootstrap(app application.App) error {
	app.AddServer(server.NewMetricsRunner(
		app.Config().MetricsPort,
		app.Logger(),
		app.PrometheusRegistry()),
	)

	return nil
}
