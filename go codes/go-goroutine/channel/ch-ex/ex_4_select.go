package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func(mesg string) {
		ch1 <- mesg
	}("Bitch Baby")

	go func(num int) {
		for i := 0; i < num; i++ {
			ch2 <- i
		}

	}(10)

	select {
	case value1 := <-ch1:
		// time.Sleep(250 * time.Millisecond)
		fmt.Println("Value1 :", value1)
	case value2 := <-ch2:
		for i := 0; i <= 10; i++ {
			time.Sleep(400 * time.Millisecond)
			fmt.Println("Value2 :", value2)
		}

	}
}
