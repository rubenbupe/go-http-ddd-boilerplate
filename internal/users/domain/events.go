package domain

import (
	"github.com/rubenbupe/go-auth-server/kit/event"
)

const UserCreatedEventType event.Type = "events.user.created"

type UserCreatedEvent struct {
	event.BaseEvent
	id   string
	name string
}

func NewUserCreatedEvent(id, name string) UserCreatedEvent {
	return UserCreatedEvent{
		id:   id,
		name: name,

		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e UserCreatedEvent) Type() event.Type {
	return UserCreatedEventType
}

func (e UserCreatedEvent) UserID() string {
	return e.id
}

func (e UserCreatedEvent) UserName() string {
	return e.name
}
