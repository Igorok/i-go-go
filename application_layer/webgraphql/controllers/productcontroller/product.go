package productcontroller

import (
	"delivery-go/domain_layer/productpkg"
	"delivery-go/application_layer/webgraphql/webtypes/producttype"

	"github.com/graphql-go/graphql"
)

// ProductController - web controller
type ProductController struct {
	productcase productpkg.UseCase
}

// NewProductController - constructor for ProductController
func NewProductController(productcase productpkg.UseCase) *ProductController {
	return &ProductController{
		productcase: productcase,
	}
}

// Products - products list
func (controller ProductController) Products() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(producttype.ProductType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return controller.productcase.GetProducts(p.Context)
		},
	}
}

// Product - product detail
func (controller ProductController) Product() *graphql.Field {
	return &graphql.Field{
		Type: producttype.ProductType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(string)
			if !ok {
				return nil, productpkg.ErrProductNotFound
			}
			return controller.productcase.GetProduct(p.Context, id)
		},
	}
}
