package clientpkg

import (
	"context"
	"delivery-go/entities_layer/user/userentity"
)

type UseCase interface {
	CreateClient(ctx context.Context, customer *userentity.UserSystem, name, email, phoneNumber string) error
}
