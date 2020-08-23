package webgraphql

import (
	"i-go-go/domain_layer/orderpkg"
	"i-go-go/domain_layer/productpkg"
	"i-go-go/domain_layer/userpkg"
	"i-go-go/application_layer/webgraphql/controllers/ordercontroller"
	"i-go-go/application_layer/webgraphql/controllers/productcontroller"
	"i-go-go/application_layer/webgraphql/controllers/usercontroller"

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
