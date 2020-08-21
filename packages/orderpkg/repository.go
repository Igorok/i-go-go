package orderpkg

import (
	"context"
	"delivery-go/entities/order/orderentity"
	"delivery-go/entities/product/productentity"
	"delivery-go/entities/user/userentity"
)

// Repository - of database methods
type Repository interface {
	GetOrder(ctx context.Context, user *userentity.UserSystem, id string) (*orderentity.Order, error)
	GetOrders(ctx context.Context, user *userentity.UserSystem, skip, limit int) ([]*orderentity.Order, error)
	CreateOrder(ctx context.Context, order *orderentity.Order) (*orderentity.Order, error)
	UpdateStatus(
		ctx context.Context,
		user *userentity.UserSystem,
		id, status string,
	) (*orderentity.Order, error)
	GetProducts(ctx context.Context, ids []string) ([]*productentity.Product, error)

	InsertLogs(ctx context.Context, orderLog []*orderentity.OrderLog) ([]*orderentity.OrderLog, error)
}
