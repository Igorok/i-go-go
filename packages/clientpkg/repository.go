package clientpkg

import (
	"context"
	"delivery-go/entities/client/cliententity"
	"delivery-go/entities/user/userentity"
)

type Repository interface {
	CreateClient(ctx context.Context, user *userentity.UserSystem, client *cliententity.Client) error
}
