package productcase

import (
	"context"
	"delivery-go/entities_layer/product/productentity"
	"delivery-go/domain_layer/productpkg"
)

type ProductCase struct {
	productRepo productpkg.Repository
}

func NewProductCase(productRepo productpkg.Repository) *ProductCase {
	return &ProductCase{
		productRepo: productRepo,
	}
}

func (pc ProductCase) GetProducts(ctx context.Context) ([]*productentity.Product, error) {
	return pc.productRepo.GetProducts(ctx)
}

func (pc ProductCase) GetProduct(ctx context.Context, id string) (productentity.Product, error) {
	if id == "" {
		return productentity.Product{}, productpkg.ErrProductNotFound
	}
	return pc.productRepo.GetProduct(ctx, id)
}
