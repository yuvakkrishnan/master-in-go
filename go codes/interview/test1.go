package main

import (
	"fmt"
)

func main() {
	n1 := []int{10, 20, 30, 40}
	n1 = append(n1, 100)
	fmt.Println(len(n1), cap(n1))
}
