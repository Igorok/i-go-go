package userpkg

import (
	"context"
	"delivery-go/entities/user/userentity"
)

type UseCase interface {
	// SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*userentity.UserSystem, error)
}
