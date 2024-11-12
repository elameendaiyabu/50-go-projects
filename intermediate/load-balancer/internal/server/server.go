package server

import (
	"fmt"
	"go-projects/intermediate/load-balancer/internal/config"
	"time"
)

func Run(c config.Configuration) {
	numJobs := 2
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for i := 0; i < len(c.Resources); i++ {
		go worker(i, jobs, results, c.Resources)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for k := 1; k <= numJobs; k++ {
		<-results
	}
}

func worker(
	id int,
	jobs <-chan int,
	results chan<- string,
	resource []config.Resources,
) {
	for j := range jobs {
		fmt.Println("request", j, "started on server", resource[id].Server)
		time.Sleep(1 * time.Second)
		fmt.Println("request", j, "finished on server", resource[id].Server)
		results <- "resource"
	}
}
