package main

import (
	"fmt"
	"time"
)

func worker(id int, name <-chan string, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("%s worker %d started job %d\n", <-name, id, j)
		time.Sleep(time.Second)
		fmt.Printf("%s worker %d finished job %d\n", <-name, id, j)
		result <- j * 2
	}
}

func main() {
	numJobs := 5
	names := make(chan string, 3)
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	workerNames := []string{"Alice", "Bob", "Charlie"}

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, names, jobs, results)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// Send worker names to workers
	for _, name := range workerNames {
		names <- name
	}

	close(names)

	// Collect results
	for k := 1; k <= numJobs; k++ {
		fmt.Println("result", k, ": ", <-results)
	}
}
