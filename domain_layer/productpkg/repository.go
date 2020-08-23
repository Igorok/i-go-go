package productpkg

import (
	"context"
	"i-go-go/entities_layer/product/productentity"
)

type Repository interface {
	GetProducts(ctx context.Context) ([]*productentity.Product, error)
	GetProduct(ctx context.Context, id string) (productentity.Product, error)
}
