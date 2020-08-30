package service_layer

import (
	"os"

	"github.com/spf13/viper"
)

func CfgInit() error {
	cfgPath := "./service_layer"
	if os.Getenv("CFG_PATH") != "" {
		cfgPath = os.Getenv("CFG_PATH")
	}

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(cfgPath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	MongoDb := os.Getenv("MONGO_DB")
	if MongoDb != "" {
		viper.Set("mongo.db", MongoDb)
	}
	MongoHost := os.Getenv("MONGO_HOST")
	if MongoHost != "" {
		viper.Set("mongo.host", MongoHost)
	}
	MongoPort := os.Getenv("MONGO_PORT")
	if MongoPort != "" {
		viper.Set("mongo.port", MongoPort)
	}
	MongoUser := os.Getenv("MONGO_USER")
	if MongoUser != "" {
		viper.Set("mongo.user", MongoUser)
	}
	MongoPassword := os.Getenv("MONGO_PASSWORD")
	if MongoPassword != "" {
		viper.Set("mongo.password", MongoPassword)
	}
	MongoAuth := os.Getenv("MONGO_AUTH")
	if MongoAuth != "" {
		viper.Set("mongo.auth", MongoAuth)
	}
	MongoReplicaset := os.Getenv("MONGO_REPLICASET")
	if MongoReplicaset != "" {
		viper.Set("mongo.replicaset", MongoReplicaset)
	}

	AppPort := os.Getenv("APP_PORT")
	if AppPort != "" {
		viper.Set("app.port", AppPort)
	}

	return nil
}
