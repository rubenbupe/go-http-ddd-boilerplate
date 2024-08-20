package create

import (
	"context"

	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
	"github.com/rubenbupe/go-auth-server/kit/event"
)

type UserService struct {
	userRepository usersdomain.UserRepository
	eventBus       event.Bus
}

func NewUserService(userRepository usersdomain.UserRepository, eventBus event.Bus) UserService {
	return UserService{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

func (s UserService) CreateUser(ctx context.Context, id, name string) error {
	userID, err := usersdomain.NewUserID(id)
	if err != nil {
		return err
	}

	userExists, err := s.userRepository.Exists(ctx, userID)

	if err != nil {
		return err
	}

	if userExists {
		return usersdomain.ErrUserAlreadyExists
	}

	user, err := usersdomain.NewUser(id, name)
	if err != nil {
		return err
	}

	if err := s.userRepository.Save(ctx, user); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, user.PullEvents())
}
