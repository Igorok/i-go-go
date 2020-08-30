package service_layer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInit connection to mongodb
func MongoInit(dbname string) *mongo.Database {
	if dbname == "" {
		dbname = viper.GetString("mongo.db")
	}

	// connection string
	uri := "mongodb://"
	if viper.GetBool("mongo.auth") {
		uri += viper.GetString("mongo.user") +
			":" +
			viper.GetString("mongo.password") +
			"@"
	}
	uri += viper.GetString("mongo.host") +
		":" +
		viper.GetString("mongo.port")

	replicaSet := viper.GetString("mongo.replicaset")
	if replicaSet != "" {
		uri += "/?replicaSet=" + replicaSet + "&authSource=admin"
	}

	fmt.Println("uri", uri)

	// create connection
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(dbname)
}
