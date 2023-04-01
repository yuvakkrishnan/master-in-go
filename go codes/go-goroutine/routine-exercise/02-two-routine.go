package main

import (
	"fmt"
	"sync"
)

func sender(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("Received", v)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)

	go sender(ch, &wg)
	go receiver(ch, &wg)

	wg.Wait()
	close(ch)
}
