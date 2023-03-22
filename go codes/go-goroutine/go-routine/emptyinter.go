package main

import (
	"fmt"
)

func main() {
	// Create a channel of interface type
	c := make(chan interface{})

	// Start a goroutine to send values to the channel
	go func() {
		c <- "hello"
		c <- 42
		c <- true
		close(c)
	}()

	// Pass the channel to a function to print the values
	printValues(c)
}

func printValues(c chan interface{}) {
	// Receive values from the channel using a for loop
	for v := range c {
		fmt.Println(v)
	}
}
