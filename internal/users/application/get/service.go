package get

import (
	"context"

	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
)

type UserService struct {
	userRepository usersdomain.UserRepository
}

func NewUserService(userRepository usersdomain.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) GetUser(ctx context.Context, id string) (*usersdomain.User, error) {
	userID, err := usersdomain.NewUserID(id)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
