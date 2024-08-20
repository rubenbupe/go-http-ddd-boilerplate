package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rubenbupe/go-auth-server/kit/event"
)

var ErrInvalidUserID = errors.New("invalid User ID")

type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return UserID{
		value: v.String(),
	}, nil
}

func (id UserID) String() string {
	return id.value
}

var ErrEmptyUserName = errors.New("the field User Name can not be empty")

type UserName struct {
	value string
}

func NewUserName(value string) (UserName, error) {
	if value == "" {
		return UserName{}, ErrEmptyUserName
	}

	return UserName{
		value: value,
	}, nil
}

func (name UserName) String() string {
	return name.value
}

var ErrUserAlreadyExists = errors.New("user already exists")

type User struct {
	Id   UserID
	Name UserName

	events []event.Event
}

type UserRepository interface {
	Save(ctx context.Context, user User) error
	Exists(ctx context.Context, id UserID) (bool, error)
	Get(ctx context.Context, id UserID) (*User, error)
}

//mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=UserRepository

func NewUser(id, name string) (User, error) {
	idVO, err := NewUserID(id)
	if err != nil {
		return User{}, err
	}

	nameVO, err := NewUserName(name)
	if err != nil {
		return User{}, err
	}

	user := User{
		Id:   idVO,
		Name: nameVO,
	}

	user.Record(NewUserCreatedEvent(idVO.String(), nameVO.String()))
	return user, nil
}

func (c *User) Record(evt event.Event) {
	c.events = append(c.events, evt)
}

func (c User) PullEvents() []event.Event {
	evt := c.events
	c.events = []event.Event{}

	return evt
}
