package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	value := <-c
	fmt.Println(value)
}
