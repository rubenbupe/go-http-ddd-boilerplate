package get

import (
	"context"
	"errors"

	"github.com/rubenbupe/go-auth-server/kit/query"
)

const UserQueryType query.Type = "query.user.get"

type UserQuery struct {
	id string
}

func NewUserQuery(id string) UserQuery {
	return UserQuery{
		id: id,
	}
}

func (c UserQuery) Type() query.Type {
	return UserQueryType
}

type UserQueryHandler struct {
	service UserService
}

func NewUserQueryHandler(service UserService) UserQueryHandler {
	return UserQueryHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h UserQueryHandler) Handle(ctx context.Context, cmd query.Query) (interface{}, error) {
	createUserCmd, ok := cmd.(UserQuery)
	if !ok {
		return nil, errors.New("unexpected query")
	}

	return h.service.GetUser(
		ctx,
		createUserCmd.id,
	)
}

func (h UserQueryHandler) SubscribedTo() query.Type {
	return UserQueryType
}
