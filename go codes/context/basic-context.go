package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create a new context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// start a goroutine to do some work
	go doWork(ctx)

	// wait for the context to be cancelled or timeout
	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled or timed out")
	}
}

func doWork(ctx context.Context) {
	// simulate some work taking 5 seconds to complete
	time.Sleep(5 * time.Second)

	// check if the context has been cancelled
	if ctx.Err() != nil {
		fmt.Println("Work cancelled")
		return
	}

	// work completed successfully
	fmt.Println("Work completed")
}
