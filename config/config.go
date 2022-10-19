package config

import (
	"log"

	"github.com/spf13/viper"
)

var config AppConfig

type AppConfig struct {
	ServerPort int `mapstructure:"SERVER_PORT"`
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error when reading config file: ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Error when unmashaling config: ", err)
	}
}

func ServerPort() int {
	return config.ServerPort
}
