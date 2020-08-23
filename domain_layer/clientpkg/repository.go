package clientpkg

import (
	"context"
	"delivery-go/entities_layer/client/cliententity"
	"delivery-go/entities_layer/user/userentity"
)

type Repository interface {
	CreateClient(ctx context.Context, user *userentity.UserSystem, client *cliententity.Client) error
}
