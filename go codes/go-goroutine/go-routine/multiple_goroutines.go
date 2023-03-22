// Starting multiple Goroutines
// Program to illustrate multiple goroutines

package main

import (
	"fmt"
	"time"
)

func dispaly(value string) {

	fmt.Println("Hello", value)

}

func main() {
	go dispaly("Yuvak")
	go dispaly("Karish")
	go dispaly("Dart")
	time.Sleep(time.Second * 1)

	fmt.Println("Main end")

}

// Benefits of Goroutines
// Here are some of the major benefits of goroutines.

// With Goroutines, concurrency is achieved in Go programming. It helps two or more independent functions to run together.
// Goroutines can be used to run background operations in a program.
// It communicates through private channels so the communication between them is safer.
// With goroutines, we can split one task into different segments to perform better.
