package config

import (
	"log"

	"github.com/spf13/viper"
)

type resources struct {
	Id     string
	Server string
}

type configuration struct {
	Endpoint struct {
		Host string
		Port int
	}
	Resources []resources
}

var Config *configuration

func ParseYAML() *configuration {
	viper.AddConfigPath(
		"/home/alamin/Development/50-go-projects/intermediate/load-balancer/internal/data/",
	)
	viper.SetConfigName("balancer.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Println("couldnt unmarshal yaml file")
	}

	return Config
}
