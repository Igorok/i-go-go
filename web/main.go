package main

import (
	"delivery-go/packages/orderpkg"
	"delivery-go/packages/orderpkg/ordercase"
	"delivery-go/packages/orderpkg/repository/ordermongo"
	"delivery-go/packages/productpkg"
	"delivery-go/packages/productpkg/productcase"
	"delivery-go/packages/productpkg/repository/productmongo"
	"delivery-go/packages/userpkg"
	"delivery-go/packages/userpkg/repository/usermongo"
	"delivery-go/packages/userpkg/usercase"
	"delivery-go/utils"
	"delivery-go/web/webgraphql"
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
	if err := utils.CfgInit(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	mongoDb := utils.MongoInit("")
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
