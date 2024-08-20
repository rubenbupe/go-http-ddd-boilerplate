package app

import (
	statushandlers "github.com/rubenbupe/go-auth-server/internal/status/platform/server/handler"
	usershandlers "github.com/rubenbupe/go-auth-server/internal/users/platform/server/handler"
	"github.com/rubenbupe/go-auth-server/kit/command"
	"github.com/rubenbupe/go-auth-server/kit/query"
	"github.com/sarulabs/di/v2"
)

var Defs = []di.Def{
	// STATUS
	di.Def{
		Name: "status.infrastructure.controller.check",
		Build: func(ctn di.Container) (interface{}, error) {
			return statushandlers.CheckHandler(), nil
		},
	},
	// USERS
	di.Def{
		Name: "users.infrastructure.controller.create",
		Build: func(ctn di.Container) (interface{}, error) {
			commandBus := ctn.Get("shared.domain.commandbus").(command.Bus)
			return usershandlers.CreateHandler(commandBus), nil
		},
	},
	di.Def{
		Name: "users.infrastructure.controller.get",
		Build: func(ctn di.Container) (interface{}, error) {
			queryBus := ctn.Get("shared.domain.querybus").(query.Bus)
			return usershandlers.GetHandler(queryBus), nil
		},
	},
}
