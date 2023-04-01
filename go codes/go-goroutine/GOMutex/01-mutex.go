// Mutex
// A mutex, short for "mutual exclusion", is a synchronization primitive that provides exclusive access to a shared resource.
// In Go, the sync package provides a Mutex type for this purpose.
// Here's an example that demonstrates how to use a mutex to protect a shared variable:

package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func incrFoo() {
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrFoo()
		}()
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}
