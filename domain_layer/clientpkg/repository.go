package clientpkg

import (
	"context"
	"i-go-go/entities_layer/client/cliententity"
	"i-go-go/entities_layer/user/userentity"
)

type Repository interface {
	CreateClient(ctx context.Context, user *userentity.UserSystem, client *cliententity.Client) error
}
