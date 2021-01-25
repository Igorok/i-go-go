package main

import (
	"i-go-go/application_layer/webgraphql"
	"i-go-go/domain_layer/orderpkg"
	"i-go-go/domain_layer/orderpkg/ordercase"
	"i-go-go/domain_layer/orderpkg/repository/ordermongo"
	"i-go-go/domain_layer/productpkg"
	"i-go-go/domain_layer/productpkg/productcase"
	"i-go-go/domain_layer/productpkg/repository/productmongo"
	"i-go-go/domain_layer/userpkg"
	"i-go-go/domain_layer/userpkg/repository/usermongo"
	"i-go-go/domain_layer/userpkg/usercase"
	"i-go-go/service_layer"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/spf13/viper"
	ginprometheus "github.com/zsais/go-gin-prometheus"
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

	r := gin.Default()

	// prometheus
	p := ginprometheus.NewPrometheus("gin")
	// If you have for instance a /customer/:name
	// p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
	// 	url := c.Request.URL.Path
	// 	for _, p := range c.Params {
	// 		if p.Key == "name" {
	// 			url = strings.Replace(url, p.Value, ":name", 1)
	// 			break
	// 		}
	// 	}
	// 	return url
	// }
	p.Use(r)

	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	graphHandler := getHandler(mongoDb)
	r.GET("/graphql", graphHandler)
	r.POST("/graphql", graphHandler)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":" + viper.GetString("app.port"))
}
