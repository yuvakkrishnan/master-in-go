// Context Errors
// The context object also features Err() which can be useful for when you need to return the error that caused your fucntion to halt.

// Whenever you call this Err() function, it’s going to return nil if Done is not yet closed. If Done is closed, then Err is going to return the error explaining why it was closed.

// In the below case, we’ll see that this returns with a context deadline exceeded error:
package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingCool(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	fmt.Println("Go Context Tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
	}

	time.Sleep(2 * time.Second)
}
