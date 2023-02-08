package main

import "fmt"

type rectangle struct {
	length int
	bredth float64
	color  string
}

func main() {

	fmt.Println(rectangle{10, 2.66, "Red"})
}
