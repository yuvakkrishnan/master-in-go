In Go, channels can be either buffered or unbuffered.

An unbuffered channel is a channel where the sender blocks until the receiver is ready to receive the message. This means that the sender will wait until there is a receiver available to receive the message before it continues execution. An unbuffered channel is created using the make function with no arguments, like this:

c := make(chan int)

Here's an example of how an unbuffered channel can be used to synchronize two Goroutines:

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

In this example, the sum function is executed as a separate Goroutine. The Goroutine calculates the sum of a slice of integers and sends the result to the channel c. The main Goroutine splits the slice into two parts and starts two Goroutines, each of which calculates the sum of one of the parts. Finally, the main Goroutine receives the results from the Goroutines using the <- operator and adds them together.

A buffered channel is a channel that can store a specified number of messages before blocking