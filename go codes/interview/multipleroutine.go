// GR1 iterates over a loop and pushes numbers from 1 to 1 million and pushes to two channels say C1.1 & C1.2 alternatively
// There are a configurable number of GRs(say GR 2.1 to GR 2.x) reading these numbers from channels C1.1 & C1.2, printing them in console and pushing to another channel say C2
// Then GR1 should read the incoming numbers from channel C2 and print a “processed <number>” statement in console
// The application should not exit until all the GRs have completed execution and exited

package main

import (
	"fmt"
	"sync"
)

func gr1(c1, c2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000000; i++ {
		if i%2 == 1 {
			c1 <- i
		} else {
			c1 <- i
		}
	}
	close(c1)
	for n := range c2 {
		fmt.Printf("processed %d\n", n)
	}

}

func gr2(id int, c1, c2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range c1 {
		fmt.Printf("GR%d: %d\n", id, n)
		c2 <- n
	}

}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go gr1(c1, c2, &wg)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go gr2(i, c1, c2, &wg)
	}
	wg.Wait()

}
