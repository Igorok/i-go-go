package service_layer

import "github.com/spf13/viper"

func CfgInit() error {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./service_layer")

	return viper.ReadInConfig()
}
