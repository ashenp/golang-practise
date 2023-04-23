package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile("server.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
