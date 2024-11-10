package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type url struct {
	Url string
}

type endpoints struct {
	Urls []url
}

var DestinationUrl *endpoints

// getting urls stored in a yaml file
func getURLs() *endpoints {
	viper.AddConfigPath(".")
	viper.SetConfigName("urls")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&DestinationUrl); err != nil {
		log.Fatal(err)
	}

	return DestinationUrl
}

func main() {
	// getting the urls
	dest := getURLs()
	fmt.Println(dest)
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
