package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf(" %d ", i)
	}
}

func words() {
	vovels := []string{"a", "e", "i", "o", "u"}
	for _, v := range vovels {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf(" %s ", v)
	}
}

func main() {

	go numbers()
	go words()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main Terminates !!! ")

}
