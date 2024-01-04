package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {

	fmt.Println("Initializing config.....")

	config = viper.New()

	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath("config/")

	err := config.ReadInConfig()

	if err != nil {

		fmt.Println("There was an exception", err)
		panic("Loading of config failed...")
	}
}

func GetConfig() *viper.Viper {

	return config
}
