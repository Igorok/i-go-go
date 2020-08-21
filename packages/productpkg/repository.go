package productpkg

import (
	"context"
	"delivery-go/entities/product/productentity"
)

type Repository interface {
	GetProducts(ctx context.Context) ([]*productentity.Product, error)
	GetProduct(ctx context.Context, id string) (productentity.Product, error)
}
