package bootstrappers

import (
	"github.com/DuC-cnZj/mars/pkg/contracts"
	"github.com/DuC-cnZj/mars/pkg/event"
	"github.com/DuC-cnZj/mars/pkg/mlog"
)

type EventBootstrapper struct{}

var events map[contracts.Event][]contracts.Listener = map[contracts.Event][]contracts.Listener{}

func (e *EventBootstrapper) Bootstrap(app contracts.ApplicationInterface) error {
	app.SetEventDispatcher(event.NewDispatcher(app))

	for e, listeners := range events {
		for _, listener := range listeners {
			app.EventDispatcher().Listen(e, listener)
		}
	}

	mlog.Debug("EventBootstrapper booted.")

	return nil
}
