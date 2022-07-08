package db

import (
	"firstProject/model"
	"github.com/spf13/viper"
)

func LoadAppConfig() string {
	var appConfig model.Config
	println("Loading server configurations")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		print("Failed to read file:", err.Error())
	}

	if err := viper.Unmarshal(&appConfig); err != nil {
		print("Failed to load file:", err.Error())
	}

	return appConfig.String()
}
