package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")

}
func main() {
	go hello()
	// time.Sleep(time.Second * 1) Run without using this line to better understand...
	fmt.Println("main function")
}

// In line no. 11, go hello() starts a new Goroutine.
// Now the hello() function will run concurrently along with the main() function.
// The main function runs in its own Goroutine and it's called the main Goroutine.
// Run this program and you will have a surprise!

// This program only outputs the text main function.
// What happened to the Goroutine we started? We need to understand the two main properties of goroutines to understand why this happens.

// A goroutine in Go is a lightweight thread of execution.
// When you start a new goroutine, the control returns immediately to the next line of code, without waiting for the goroutine to finish.
// The goroutine runs concurrently with the rest of your program, and its return values are ignored.

// Think of a goroutine as a separate task that runs alongside your main program.
// When you start a goroutine, it runs independently, and you don't have to wait for it to finish before you can move on to the next task.
// The main goroutine is the first one that is started when you run your program.
// It is the "main" task, and any other goroutines are "child" tasks.
// If the main goroutine finishes, the program will terminate and all other goroutines will be terminated as well. So, it's important to keep the main goroutine running for as long as you want the program to continue executing.
