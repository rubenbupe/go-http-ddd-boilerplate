package shared

import (
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/bus/inmemory"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/storage"
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/storage/postgres"
	"github.com/sarulabs/di/v2"
)

var Defs = []di.Def{
	// BUSES
	di.Def{
		Name: "shared.domain.commandbus",
		Build: func(ctn di.Container) (interface{}, error) {
			return inmemory.NewCommandBus(), nil
		},
	},
	di.Def{
		Name: "shared.domain.querybus",
		Build: func(ctn di.Container) (interface{}, error) {
			return inmemory.NewQueryBus(), nil
		},
	},
	di.Def{
		Name: "shared.domain.eventbus",
		Build: func(ctn di.Container) (interface{}, error) {
			return inmemory.NewEventBus(), nil
		},
	},

	// DB
	di.Def{
		Name: "shared.infrastructure.sqlconfig",
		Build: func(ctn di.Container) (interface{}, error) {
			return postgres.CreateConfig()
		},
	},
	di.Def{
		Name: "shared.infrastructure.sqlconnection",
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("shared.infrastructure.sqlconfig").(*storage.Dbconfig)
			return storage.CreateConnection("shared", cfg)
		},
	},
}
