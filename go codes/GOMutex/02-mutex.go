package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func increment() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()
	fmt.Println("Counter:", counter)
}
