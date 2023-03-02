// Program to run two goroutines concurrently

package main

import (
	"fmt"
	"time"
)

// create a function
func display(message string) {

	// infinite for loop
	for {

		fmt.Println(message)
		// to sleep the process for 1 second
		time.Sleep(time.Second * 1)

	}
}

func main() {

	// call function with Child Thread goroutine
	go display("Process 1")
	//Main Thread
	display("Process 2")

}

// Run two functions concurrently using Goroutine
// In a concurrent program, the main() is always a default goroutine.
// Other goroutines can not execute if the main() is not executing.
// So, in order to make sure that all the goroutines are executed before the main function ends, we sleep the process so that the other processes get a chance to execute.

// In the above example, we have added a line in the function definition.

// time.Sleep(time.Second * 1)

// Here, when the display("Process 2") is executing, the time.Sleep() function stops the process for 1 second. In that 1 second, the goroutine go display("Process 1") is executed.

// This way, the functions run concurrently before the main() functions stops.
