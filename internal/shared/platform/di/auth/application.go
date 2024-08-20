package auth

import (
	"github.com/rubenbupe/go-auth-server/internal/shared/platform/storage"
	"github.com/rubenbupe/go-auth-server/kit/event"
	"github.com/sarulabs/di/v2"

	usercreate "github.com/rubenbupe/go-auth-server/internal/users/application/create"
	userget "github.com/rubenbupe/go-auth-server/internal/users/application/get"
	userincrease "github.com/rubenbupe/go-auth-server/internal/users/application/increase"
	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
	userspostgres "github.com/rubenbupe/go-auth-server/internal/users/platform/storage/postgres"
)

var Defs = []di.Def{
	// REPOSITORIES
	di.Def{
		Name: "users.domain.repository",
		Build: func(ctn di.Container) (interface{}, error) {
			conn := ctn.Get("shared.infrastructure.sqlconnection").(*storage.Connection)
			dbconfig := ctn.Get("shared.infrastructure.sqlconfig").(*storage.Dbconfig)
			return userspostgres.NewUserRepository(conn, dbconfig), nil
		},
	},
	// USE CASES, COMMAND HANDLERS, AND EVENT HANDLERS
	di.Def{
		Name: "users.domain.create",
		Build: func(ctn di.Container) (interface{}, error) {
			repo := ctn.Get("users.domain.repository").(usersdomain.UserRepository)
			eventBus := ctn.Get("shared.domain.eventbus").(event.Bus)
			return usercreate.NewUserService(repo, eventBus), nil
		},
	},
	di.Def{
		Name: "users.domain.createcommandhandler",
		Build: func(ctn di.Container) (interface{}, error) {
			service := ctn.Get("users.domain.create").(usercreate.UserService)
			return usercreate.NewUserCommandHandler(service), nil
		},
		Tags: []di.Tag{
			{Name: "command-handler"},
		},
	},
	di.Def{
		Name: "users.domain.increment",
		Build: func(ctn di.Container) (interface{}, error) {
			return userincrease.NewUserCounterService(), nil
		},
	},
	di.Def{
		Name: "users.domain.incrementonusercreated",
		Build: func(ctn di.Container) (interface{}, error) {
			service := ctn.Get("users.domain.increment").(userincrease.UserCounterService)
			return userincrease.NewIncreaseUsersCounterOnUserCreated(service), nil
		},
		Tags: []di.Tag{
			{Name: "event-handler"},
		},
	},
	di.Def{
		Name: "users.domain.get",
		Build: func(ctn di.Container) (interface{}, error) {
			repo := ctn.Get("users.domain.repository").(usersdomain.UserRepository)
			return userget.NewUserService(repo), nil
		},
	},
	di.Def{
		Name: "users.domain.getqueryhandler",
		Build: func(ctn di.Container) (interface{}, error) {
			service := ctn.Get("users.domain.get").(userget.UserService)
			return userget.NewUserQueryHandler(service), nil
		},
		Tags: []di.Tag{
			{Name: "query-handler"},
		},
	},
}
