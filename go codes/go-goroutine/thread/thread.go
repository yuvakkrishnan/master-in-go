package main

import (
	"fmt"
	"sync"
)

func waiterGrp(id int, wg *sync.WaitGroup) {
	fmt.Printf("Id Worker %d \n", id)
	sum := 0
	for i := 1; i <= id; i++ {
		sum += i * 100
	}
	fmt.Printf("Worker ID : %d workers Salary %d\n", id, sum)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go waiterGrp(i, &wg)
	}
	wg.Wait()

	fmt.Println("All workers done... The main Thread Ends")
}
