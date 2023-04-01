package main

import (
	"fmt"
	"sync"
)

func fooIncr(values *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		*values++
	}

}

func main() {
	var value int = 0
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go fooIncr(&value, &wg)
	}
	wg.Wait()
	fmt.Println("Final value:value", value)

}
