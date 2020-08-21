package orderpkg

import (
	"context"
	"delivery-go/entities/order/orderentity"
	"delivery-go/entities/user/userentity"
	"time"
)

// UseCase - interface for order package
type UseCase interface {
	GetOrder(ctx context.Context, user *userentity.UserSystem, id string) (*orderentity.Order, error)
	GetOrders(ctx context.Context, user *userentity.UserSystem, skip, limit int) ([]*orderentity.Order, error)
	CreateOrder(
		ctx context.Context,
		user *userentity.UserSystem,
		cart []orderentity.CartItem,
		dateCustomer time.Time,
		address string,
		comment string,
	) (*orderentity.Order, error)
	UpdateStatus(ctx context.Context, user *userentity.UserSystem, status string) (*orderentity.Order, error)
}
