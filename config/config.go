package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config AppConfig

type ServerConfig struct {
	Port int `yaml:"port"`
}

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error when reading config file: ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Error when unmashaling config: ", err)
	}

	fmt.Printf("%x", config)
}

func ServerPort() int {
	return config.Server.Port
}
