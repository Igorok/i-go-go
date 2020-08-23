package productmongo_test

import (
	"context"
	"delivery-go/domain_layer/productpkg/repository/productmongo"
	"delivery-go/service_layer"
	"os"

	"delivery-go/bin/testdata"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbName string = "delivery_test_products"
const prCollName string = "test_products"

var db *mongo.Database

func TestMain(m *testing.M) {

	// You create an Person and you save in database
	setUp()
	retCode := m.Run()

	// When you have executed the test, the Person is deleted from database
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("../../../../service_layer")
	viper.ReadInConfig()

	db = service_layer.MongoInit(dbName)
	testdata.InsertProducts(db, prCollName)
}

func tearDown() {
	db.Drop(context.TODO())
}

func TestGetProducts(t *testing.T) {
	t.Log("Test product list repository")

	productRepo := productmongo.NewProductRepository(db, prCollName)
	pmArr, err := productRepo.GetProducts(context.TODO())

	assert.Nil(t, err)
	assert.NotNil(t, pmArr)
	assert.NotEqual(t, len(pmArr), 0)
}

func TestGetProduct(t *testing.T) {
	t.Log("Test product detail repository")

	productRepo := productmongo.NewProductRepository(db, prCollName)
	pm, err := productRepo.GetProduct(context.TODO(), "5e874f4b327272d07e537a4d")

	assert.Nil(t, err)
	assert.NotNil(t, pm)
	assert.Equal(t, pm.Price, 200)
	assert.Equal(t, pm.Name, "Steak New York")
}
