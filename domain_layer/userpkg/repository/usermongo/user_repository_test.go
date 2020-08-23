package usermongo_test

import (
	"context"
	"delivery-go/bin/testdata"
	"delivery-go/domain_layer/userpkg/repository/usermongo"
	"delivery-go/service_layer"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbName string = "delivery_test_users_system"
const usrCollName string = "test_users_system"

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
	testdata.InsertUsers(db, usrCollName)
}

func tearDown() {
	db.Drop(context.TODO())
}

func TestGetUser(t *testing.T) {
	t.Log("Test user detail repository")

	userRepo := usermongo.NewUserRepository(db, usrCollName)
	um, err := userRepo.GetUser(context.TODO(), "courier_pizza", "a4a165fc8fca19f70e88fad37d5476ba1d0b7415")

	assert.Nil(t, err)
	assert.NotNil(t, um)
	assert.Equal(t, (*um).Login, "courier_pizza")
	assert.Equal(t, (*um).ID, "5e874f4b327272d07e537a50")
}
