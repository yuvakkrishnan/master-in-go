package main

import (
	"fmt"
)

func getSquareChan(input <-chan int) <-chan int {
	output := make(chan int, 100)

	go func() {
		for v := range input {
			output <- v * v
		}
		close(output)
	}()
	return output
}

func main() {

	input := make(chan int)

	output := getSquareChan(input)

	go func() {
		for i := 0; i < 5; i++ {
			input <- i
		}
		close(input)
	}()

	for v := range output {
		fmt.Println(v)
	}

}
