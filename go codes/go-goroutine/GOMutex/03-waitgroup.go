// A WaitGroup is a synchronization primitive that allows a program to wait for a collection of goroutines to complete. In Go, the sync package provides a WaitGroup type for this purpose. Here's an example that demonstrates how to use a WaitGroup:

package main

import (
	"fmt"
	"sync"
	"time"
)

func waitFoo(id int, yk *sync.WaitGroup) {
	defer yk.Done()
	fmt.Printf("Workers %d started wroking\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d ended working\n", id)
}

func main() {
	var yk sync.WaitGroup

	for i := 1; i <= 5; i++ {
		yk.Add(1)
		go waitFoo(i, &yk)
	}
	yk.Wait()

	fmt.Println("All worker done")

}
