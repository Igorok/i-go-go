package utils

import "github.com/spf13/viper"

func CfgInit() error {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./utils")

	return viper.ReadInConfig()
}
