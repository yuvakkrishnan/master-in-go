package main

import "fmt"

type rectangle struct {
	length   int
	bredth   int
	color    string
	geometry struct {
		area      int
		perimeter int
	}
}

func main() {
	var rect rectangle //Dot Notation

	rect.length = 10
	rect.bredth = 20
	rect.color = "Green"
	rect.geometry.area = rect.length * rect.bredth

	rect.geometry.perimeter = 2 * (rect.length * rect.bredth)

	fmt.Println(rect)

	fmt.Println("Rectangle of area", rect.geometry.area)

	fmt.Println("Rectangle of perimeter ", rect.geometry.perimeter)

}
