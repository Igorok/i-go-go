package productcase_test

import (
	"context"
	"testing"

	"delivery-go/packages/productpkg/productcase"
	"delivery-go/packages/productpkg/repository/productmock"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	pr := productmock.ProductRepositoryMock{}
	pc := productcase.NewProductCase(pr)

	pmArr, err := pc.GetProducts(context.TODO())

	t.Log("pmArr", pmArr)

	assert.Nil(t, err)
	assert.NotNil(t, pmArr)
	assert.NotEqual(t, len(pmArr), 0)
}

func TestGetProduct(t *testing.T) {
	pr := productmock.ProductRepositoryMock{}
	pc := productcase.NewProductCase(pr)

	pm, err := pc.GetProduct(context.TODO(), "5e874f4b327272d07e537a4d")

	t.Log("pm", pm)

	assert.Nil(t, err)
	assert.NotNil(t, pm)
	assert.Equal(t, pm.Price, 200)
	assert.Equal(t, pm.Name, "Steak New York")
}
