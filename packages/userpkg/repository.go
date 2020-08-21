package userpkg

import (
	"context"
	"delivery-go/entities/user/userentity"
)

type Repository interface {
	// CreateUser(ctx context.Context, user *userentity.UserSystem) error
	GetUser(ctx context.Context, login, password string) (*userentity.UserSystem, error)
}
