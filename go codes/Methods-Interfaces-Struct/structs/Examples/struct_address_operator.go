//struct instantiation using pointer address operator

package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	rect := &rectangle{10, 20, "White"}
	fmt.Println(rect)
	rect1 := &rectangle{}
	rect.length = 20
	rect.breadth = 30
	fmt.Println(rect1)

	rect2 := &rectangle{}
	(*rect2).breadth = 20
	(*rect2).length = 10

	fmt.Println(rect2)
}
