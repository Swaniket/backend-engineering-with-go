// This patterns shows an use case for channels & go-routine

package main

import (
	"fmt"
	"time"
)

// Here worker is just a function, but it can be a server/cron job etc.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs) // 5 clients
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ { // 3 workers for those 5 clients
		go worker(w, jobs, results) // Starting each worker in go-routine
	}

	time.Sleep(2 * time.Second)

	for j := 1; j <= numJobs; j++ {
		jobs <- j // Add the jobs into jobs channel
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
