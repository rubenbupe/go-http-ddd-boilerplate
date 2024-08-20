package create

import (
	"context"
	"errors"
	"testing"

	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
	"github.com/rubenbupe/go-auth-server/internal/users/platform/storage/storagemocks"
	"github.com/rubenbupe/go-auth-server/kit/event"
	"github.com/rubenbupe/go-auth-server/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UserService_CreateUser_RepositoryError(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Test User"

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Exists", mock.Anything, mock.AnythingOfType("domain.UserID")).Return(false, nil)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.User")).Return(errors.New("something unexpected happened"))

	eventBusMock := new(eventmocks.Bus)

	userService := NewUserService(userRepositoryMock, eventBusMock)

	err := userService.CreateUser(context.Background(), userID, userName)

	userRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UserService_CreateUser_EventsBusError(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Test User"

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Exists", mock.Anything, mock.AnythingOfType("domain.UserID")).Return(false, nil)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.User")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	userService := NewUserService(userRepositoryMock, eventBusMock)

	err := userService.CreateUser(context.Background(), userID, userName)

	userRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UserService_CreateUser_Succeed(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Test User"

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Exists", mock.Anything, mock.AnythingOfType("domain.UserID")).Return(false, nil)
	userRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.User")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.MatchedBy(func(events []event.Event) bool {
		evt := events[0].(usersdomain.UserCreatedEvent)
		return evt.UserName() == userName
	})).Return(nil)

	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	userService := NewUserService(userRepositoryMock, eventBusMock)

	err := userService.CreateUser(context.Background(), userID, userName)

	userRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_UserService_CreateUser_UserAlreadyExists(t *testing.T) {
	userID := "37a0f027-15e6-47cc-a5d2-64183281087e"
	userName := "Test User"

	userRepositoryMock := new(storagemocks.UserRepository)
	userRepositoryMock.On("Exists", mock.Anything, mock.AnythingOfType("domain.UserID")).Return(true, nil)

	eventBusMock := new(eventmocks.Bus)

	userService := NewUserService(userRepositoryMock, eventBusMock)

	err := userService.CreateUser(context.Background(), userID, userName)

	userRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
	assert.Equal(t, err, usersdomain.ErrUserAlreadyExists)
}
