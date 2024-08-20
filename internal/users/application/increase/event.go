package increase

import (
	"context"
	"errors"

	usersdomain "github.com/rubenbupe/go-auth-server/internal/users/domain"
	"github.com/rubenbupe/go-auth-server/kit/event"
)

type IncreaseUsersCounterOnUserCreated struct {
	increasingService UserCounterService
}

func NewIncreaseUsersCounterOnUserCreated(increaserService UserCounterService) IncreaseUsersCounterOnUserCreated {
	return IncreaseUsersCounterOnUserCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseUsersCounterOnUserCreated) Handle(_ context.Context, evt event.Event) error {
	userCreatedEvt, ok := evt.(usersdomain.UserCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(userCreatedEvt.ID())
}

func (e IncreaseUsersCounterOnUserCreated) SubscribedTo() event.Type {
	return usersdomain.UserCreatedEventType
}
