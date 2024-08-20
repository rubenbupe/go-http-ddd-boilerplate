package create

import (
	"context"
	"errors"

	"github.com/rubenbupe/go-auth-server/kit/command"
)

const UserCommandType command.Type = "command.user.create"

type UserCommand struct {
	id   string
	name string
}

func NewUserCommand(id, name string) UserCommand {
	return UserCommand{
		id:   id,
		name: name,
	}
}

func (c UserCommand) Type() command.Type {
	return UserCommandType
}

type UserCommandHandler struct {
	service UserService
}

func NewUserCommandHandler(service UserService) UserCommandHandler {
	return UserCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h UserCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createUserCmd, ok := cmd.(UserCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateUser(
		ctx,
		createUserCmd.id,
		createUserCmd.name,
	)
}

func (h UserCommandHandler) SubscribedTo() command.Type {
	return UserCommandType
}
