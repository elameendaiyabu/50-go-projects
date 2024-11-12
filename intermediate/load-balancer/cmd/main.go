package main

import (
	"fmt"
	"go-projects/intermediate/load-balancer/internal/config"
)

// NOTE: will have servers and an endpoint
// NOTE: when a request is made to an endpoint, it uses free servers for comm
// NOTE: STEPS: 1. to be able to load and read servers and endpoint
// NOTE: STEPS: 2. create worker pools that assign servers ...

func main() {
	fmt.Println(config.ParseYAML())
}
