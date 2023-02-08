package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	rwLock  sync.RWMutex
	mx      sync.RWMutex
)

func increment() {
	mx.Lock()
	counter++
	mx.Unlock()
}

func read(wg *sync.WaitGroup) {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Reading locking...")
	time.Sleep(time.Second)
	fmt.Println("Reading unlocking...")
	wg.Done()
}
func write(wg *sync.WaitGroup) {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write locking...")
	time.Sleep(time.Second)
	fmt.Println("Write unlocking...")
	wg.Done()
}

func readerWriter(wg *sync.WaitGroup) {
	wg.Add(5)
	go write(wg)
	go read(wg)
	go read(wg)
	go read(wg)
	go write(wg)
	time.Sleep(time.Second)
	fmt.Println("Done ...")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	readerWriter(&wg)
	wg.Wait()
	fmt.Println("Main thread is done")
}
