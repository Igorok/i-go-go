package webgraphql

import (
	"delivery-go/packages/orderpkg"
	"delivery-go/packages/productpkg"
	"delivery-go/packages/userpkg"
	"delivery-go/web/webgraphql/controllers/ordercontroller"
	"delivery-go/web/webgraphql/controllers/productcontroller"
	"delivery-go/web/webgraphql/controllers/usercontroller"

	"github.com/graphql-go/graphql"
)

// GetSchema - return graphql schema for http requests
func GetSchema(
	productcase productpkg.UseCase,
	usercase userpkg.UseCase,
	ordercase orderpkg.UseCase,
) (graphql.Schema, error) {

	userController := usercontroller.NewUserController(usercase)
	productController := productcontroller.NewProductController(productcase)
	orderController := ordercontroller.NewOrderController(ordercase, usercase)

	fields := graphql.Fields{
		"signIn":     userController.SignIn(),
		"parseToken": userController.ParseToken(),

		"products": productController.Products(),
		"product":  productController.Product(),

		"createOrder": orderController.CreateOrder(),
		"getOrders":   orderController.GetOrders(),
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	return graphql.NewSchema(schemaConfig)
}
