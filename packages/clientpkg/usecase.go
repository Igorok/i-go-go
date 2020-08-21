package clientpkg

import (
	"context"
	"delivery-go/entities/user/userentity"
)

type UseCase interface {
	CreateClient(ctx context.Context, customer *userentity.UserSystem, name, email, phoneNumber string) error
}
