package config

import (
	"log"

	"github.com/spf13/viper"
)

type Resources struct {
	Id     string
	Server string
}

type Configuration struct {
	Endpoint struct {
		Host string
		Port string
	}
	Resources []Resources
}

var Config *Configuration

func ParseYAML() *Configuration {
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
