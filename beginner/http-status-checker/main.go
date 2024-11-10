package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type url struct {
	Url string
}

type endpoints struct {
	Urls []url
}

// type for channel to grab results from http call
type result struct {
	url    string
	status string
}

var DestinationUrl *endpoints

// getting urls stored in a yaml file
func getURLs() []string {
	viper.AddConfigPath(".")
	viper.SetConfigName("urls")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&DestinationUrl); err != nil {
		log.Fatal(err)
	}

	var urls []string
	for _, url := range DestinationUrl.Urls {
		urls = append(urls, url.Url)
	}

	return urls
}

// checking urls status
func checkStatus(urls []string) {
	// creating a buffered channel
	resultChan := make(chan result, len(urls))

	// ranging over urls, making GET requests and storing data in channel
	for _, url := range urls {
		go func(u string) {
			resp, err := http.Get(u)
			if err != nil {
				log.Printf("couldnt make request for %s", u)
			}
			defer resp.Body.Close()

			resultChan <- result{url: u, status: resp.Status}
		}(url)
	}

	// looping over channel and printing result, I used select to trigger
	// second case when timeout is reached
	for i := 0; i < len(urls); i++ {
		select {
		case r := <-resultChan:
			fmt.Printf("%s \t==>\t %s\n", r.url, r.status)
		case <-time.After(5 * time.Second):
			r := <-resultChan
			fmt.Printf("%s \t==>\t %s\n", r.url, "request timeout (time > 5s)")
		}
	}
	close(resultChan)
}

func main() {
	// getting the urls
	dest := getURLs()

	checkStatus(dest)
}
