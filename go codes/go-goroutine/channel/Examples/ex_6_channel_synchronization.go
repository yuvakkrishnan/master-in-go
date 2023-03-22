package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c <- 42
	}()
	go func() {
		defer wg.Done()
		value := <-c
		fmt.Println(value)
	}()

	wg.Wait()
}
