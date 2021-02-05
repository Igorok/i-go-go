package main

import (
	"bytes"
	"encoding/json"
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
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
	"github.com/prometheus/client_golang/prometheus"
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
			viper.GetString("app.signing_key"),
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

// Metrics - middleware for prometheus metrics
func Metrics(histogram *prometheus.HistogramVec) gin.HandlerFunc {
	// GqlBody - description of grapql query
	type GqlBody struct {
		OperationName string `json:"operationName"`
		Query         string `json:"query"`
	}

	return func(c *gin.Context) {
		method := "default"
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		var gqlBody GqlBody
		gqlError := json.Unmarshal(body, &gqlBody)
		if gqlError == nil {
			method = gqlBody.OperationName
		}

		t := time.Now()

		// before request
		c.Next()
		// after request

		latency := time.Since(t).Milliseconds()

		histogram.WithLabelValues(method).Observe(float64(latency))
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
	p.Use(r)

	httpRequestHistogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "i_go_go_duration_milliseconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   []float64{1, 2, 5, 10, 20, 60},
	}, []string{"method"})
	prometheus.Register(httpRequestHistogram)

	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	graphHandler := getHandler(mongoDb)
	r.GET("/graphql", graphHandler)
	r.POST("/graphql", Metrics(httpRequestHistogram), graphHandler)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":" + viper.GetString("app.port"))
}
