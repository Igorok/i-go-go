package productcase

import (
	"context"
	"delivery-go/entities/product/productentity"
	"delivery-go/packages/productpkg"
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
