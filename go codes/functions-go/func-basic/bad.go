package main

import "fmt"

func showN() {
	if N == 0 {
		return
	}
	fmt.Printf("showN       : N is %d\n", N)
}

func incrN() {
	N++
}
