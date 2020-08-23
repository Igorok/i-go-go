package main

import (
	"delivery-go/domain_layer/orderpkg"
	"delivery-go/domain_layer/orderpkg/ordercase"
	"delivery-go/domain_layer/orderpkg/repository/ordermongo"
	"delivery-go/domain_layer/productpkg"
	"delivery-go/domain_layer/productpkg/productcase"
	"delivery-go/domain_layer/productpkg/repository/productmongo"
	"delivery-go/domain_layer/userpkg"
	"delivery-go/domain_layer/userpkg/repository/usermongo"
	"delivery-go/domain_layer/userpkg/usercase"
	"delivery-go/service_layer"
	"delivery-go/application_layer/webgraphql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func getHandler(db *mongo.Database) gin.HandlerFunc {
	type UseCases struct {
		usercase    userpkg.UseCase
		productcase productpkg.UseCase
		ordercase   orderpkg.UseCase
	}

	userRepo := usermongo.NewUserRepository(db, "users_system")
	productRepo := productmongo.NewProductRepository(db, "products")
	orderRepo := ordermongo.NewOrderRepository(db, "orders", "order_logs", "products")

	uc := UseCases{
		usercase: usercase.NewAuthUseCase(
			userRepo,
			viper.GetString("app.hash_salt"),
			[]byte(viper.GetString("app.signing_key")),
			viper.GetDuration("app.token_ttl"),
		),
		productcase: productcase.NewProductCase(productRepo),
		ordercase:   ordercase.NewOrderCase(orderRepo),
	}

	schema, _ := webgraphql.GetSchema(uc.productcase, uc.usercase, uc.ordercase)

	return func(c *gin.Context) {
		// Creates a GraphQL-go HTTP handler with the defined schema
		h := handler.New(&handler.Config{
			Schema:   &schema,
			Pretty:   true,
			GraphiQL: true,
		})

		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	if err := service_layer.CfgInit(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	mongoDb := service_layer.MongoInit("")
	graphHandler := getHandler(mongoDb)

	r := gin.Default()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	r.GET("/graphql", graphHandler)
	r.POST("/graphql", graphHandler)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":" + viper.GetString("app.port"))

}
