package get

import (
	"context"
	"errors"
	"testing"

	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
	"github.com/rubenbupe/go-auth-server/internal/users/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UserService_GetUser_RepositoryError(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Get", mock.Anything, mock.AnythingOfType("domain.UserID")).Return(((*usersdomain.User)(nil)), errors.New("something unexpected happened"))

	userService := NewUserService(userRepositoryMock)

	_, err := userService.GetUser(context.Background(), userID)

	userRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UserService_GetUser_Succeed(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Test User"

	user, err := usersdomain.NewUser(userID, userName)
	assert.NoError(t, err)

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Get", mock.Anything, mock.AnythingOfType("domain.UserID")).Return(&user, nil)

	userService := NewUserService(userRepositoryMock)

	foundUser, err := userService.GetUser(context.Background(), userID)

	assert.Equal(t, user.Id.String(), foundUser.Id.String())
	userRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
