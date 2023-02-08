// Goroutines and threads are both mechanisms for concurrency in programming, but they have some key differences:

// Creation Cost: Goroutines are much lighter weight compared to threads and are created much faster. This makes it possible to have thousands of Goroutines running concurrently in a program, whereas creating that many threads would be prohibitively expensive.

// Scheduling: In Go, Goroutines are scheduled by the Go runtime, whereas threads are typically scheduled by the operating system. This means that the Go runtime can optimize the scheduling of Goroutines to balance the workload and minimize the overhead of context switching.

// Stack Size: Goroutines have a smaller stack size compared to threads. This means that Goroutines use less memory, but also that they are less suitable for tasks that require a lot of recursion or deep call stacks.

// Synchronization: Goroutines can communicate and synchronize their execution using channels and select statements, whereas threads typically use locks, semaphores, and other synchronization mechanisms. The synchronization mechanisms in Go are designed to be simpler and safer to use compared to those in other programming languages.

// Blocking: When a Goroutine blocks, such as when it is waiting for I/O, it does not prevent other Goroutines from executing. In contrast, when a thread blocks, the entire process is blocked and other threads cannot run.

// Here is an example of using Goroutines in Go:
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, c chan int) {
	fmt.Printf("Worker %d starting\n", id)
	result := 42 * id
	c <- result
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg, c)
	}
	wg.Wait()
	close(c)
	for result := range c {
		fmt.Println("Result:", result)
	}
	fmt.Println("All workers done")
}

// In this example, the worker function is executed as a separate Goroutine using the go keyword.
//The worker function takes an id, a sync.WaitGroup, and a channel c as arguments.
//The WaitGroup is used to coordinate the execution of the Goroutines and to ensure that the main Goroutine does not exit before all the worker Goroutines are done.
//The channel c is used to communicate the result from the worker Goroutine to the main Goroutine.
// The main Goroutine creates five worker Goroutines and waits for all of them to finish using the Wait method of the WaitGroup.
//When a worker Goroutine is done, it sends the result to the channel c and calls the Done method of the WaitGroup to signal that it has finished.
// Finally, the main Goroutine receives the results from the worker Goroutines via the channel c and prints them to the console.
// This example demonstrates how Goroutines can be used to perform multiple tasks concurrently, and how channels and `WaitGroup
