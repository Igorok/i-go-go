package productmock

import (
	"context"
	"delivery-go/bin/testdata"
	"delivery-go/entities_layer/product/productentity"
)

// ProductRepositoryMock mock for productpkg.Repository
type ProductRepositoryMock struct{}

// GetProducts contain requests to get list of products
func (r ProductRepositoryMock) GetProducts(ctx context.Context) ([]*productentity.Product, error) {
	pmArr := testdata.GetTestProducts()
	pmLinksArr := make([]*productentity.Product, len(pmArr))

	for i, pmVal := range pmArr {
		pmLinksArr[i] = &pmVal
	}

	return pmLinksArr, nil
}

// GetProduct - detail of product
func (r ProductRepositoryMock) GetProduct(ctx context.Context, id string) (productentity.Product, error) {
	pmArr := testdata.GetTestProducts()
	pm := productentity.Product{}

	for _, p := range pmArr {
		if p.ID == id {
			pm = p
			break
		}
	}

	return pm, nil
}
