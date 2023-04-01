package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// create a context with a 5 second timeout
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// create a channel to receive the result
	resultChan := make(chan string)

	// start a goroutine to do some work
	go func() {
		result, err := doSomeWork()
		if err != nil {
			resultChan <- err.Error()
			return
		}
		resultChan <- result
	}()

	select {
	case <-ctxWithTimeout.Done():
		fmt.Println("Timeout exceeded")
	case result := <-resultChan:
		fmt.Println("Result:", result)
	}
}

func doSomeWork() (string, error) {
	// simulate some work taking 10 seconds to complete
	time.Sleep(10 * time.Second)
	return "done", nil
}
