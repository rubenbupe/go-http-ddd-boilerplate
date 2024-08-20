package bootstrap

import (
	"context"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/di"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/server"
	"github.com/rubenbupe/go-auth-server/kit/command"
	"github.com/rubenbupe/go-auth-server/kit/event"
	"github.com/rubenbupe/go-auth-server/kit/query"
)

func Run() error {
	var cfg config
	err := envconfig.Process("APP", &cfg)
	if err != nil {
		return err
	}

	configureCommandBus()
	configureQueryBus()
	configureEventBus()

	commandBus := di.Instance().Container.Get("shared.domain.commandbus").(command.Bus)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

func configureCommandBus() {
	diContainer := di.Instance()
	commandBus := diContainer.Container.Get("shared.domain.commandbus").(command.Bus)
	commandHandlers := diContainer.GetByTag("command-handler")

	for _, handlerDef := range commandHandlers {
		handler := diContainer.Container.Get(handlerDef).(command.Handler)
		commandBus.Register(handler.SubscribedTo(), handler)
	}
}

func configureQueryBus() {
	diContainer := di.Instance()
	queryBus := diContainer.Container.Get("shared.domain.querybus").(query.Bus)
	queryHandlers := diContainer.GetByTag("query-handler")

	for _, handlerDef := range queryHandlers {
		handler := diContainer.Container.Get(handlerDef).(query.Handler)
		queryBus.Register(handler.SubscribedTo(), handler)
	}
}

func configureEventBus() {
	diContainer := di.Instance()
	eventBus := diContainer.Container.Get("shared.domain.eventbus").(event.Bus)
	eventHandlers := diContainer.GetByTag("event-handler")

	for _, handlerDef := range eventHandlers {
		handler := diContainer.Container.Get(handlerDef).(event.Handler)
		eventBus.Subscribe(handler.SubscribedTo(), handler)
	}
}

type config struct {
	// Server configuration
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
}
