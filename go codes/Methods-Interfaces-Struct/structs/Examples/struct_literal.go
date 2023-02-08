//instance_of_struct_type using of struct literals.....

package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	var rect1 = rectangle{10, 20, "Yellow"}
	fmt.Println(rect1)
	var rect2 = rectangle{10, 5, "Yellow"} //Skipped Breath
	fmt.Println(rect2)
	var rect3 = rectangle{10, 8, "Blue"} //skiped Color
	fmt.Println(rect3)

}
