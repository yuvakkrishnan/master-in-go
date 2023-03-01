package main

import (
	"fmt"
	"sync"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}
func b() {
	for i := 'a'; i < 'o'; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go a()
	go b()
	wg.Done()
	fmt.Println("Main Routine End")
}
