// Context Deadlines using WithTimeout
// In some high-performance systems, you may need to return a response within a deadline. In a previous company I’ve worked at, we had roughly 2 seconds to return a response within our system or the action taking place would fail.

// The context.Context struct contains some of the functionality that we need in order to control how our system behaves when our system exceeds a deadline.

package main

import (
	"context"
	"fmt"
	"time"
)

func doSomeThingCool(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go context")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doSomeThingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}
	time.Sleep(2 * time.Second)
}
