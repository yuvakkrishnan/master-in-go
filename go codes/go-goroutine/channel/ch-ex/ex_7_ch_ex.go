package main

import (
	"fmt"
	"math/rand"
)

func CalVal(val chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value :{}", value)
	val <- value
}

func main() {
	fmt.Println("chan startss!!!!")
	ch := make(chan int)
	defer close(ch)

	go CalVal(ch)

	result := <-ch
	fmt.Println(result)
}
