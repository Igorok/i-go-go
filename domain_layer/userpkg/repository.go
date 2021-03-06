package userpkg

import (
	"context"
	"i-go-go/entities_layer/user/userentity"
)

type Repository interface {
	// CreateUser(ctx context.Context, user *userentity.UserSystem) error
	GetUser(ctx context.Context, login, password string) (*userentity.UserSystem, error)
}
