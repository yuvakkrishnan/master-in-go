//struct instantiation using new keyword

package main

import "fmt"

type recctangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	rect := new(recctangle)
	rect.length = 10
	rect.breadth = 20
	rect.color = "Black"

	rect2 := new(recctangle)
	rect2.length = 150
	rect2.breadth = 204
	rect2.color = "Blackmamba"

	fmt.Println("Value stored in rect = ", rect)
	fmt.Println("Address of rect = ", &rect)
	fmt.Println("Value stored in variable rect = ", rect)
	fmt.Println("Value stored in rect2 = ", rect2)
	fmt.Println("Address of rect2 = ", &rect2)
	fmt.Println("Value stored in variable rect2 = ", rect2)

}
