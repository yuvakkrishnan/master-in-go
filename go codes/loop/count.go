package main

import "fmt"

func main() {
	var x int
	if x := 100; x == 100 {
		fmt.Println("enter the number 0 to 10")
		fmt.Scanln(&x)
	} else {
		fmt.Println("try with lower number!")
	}

	// fmt.Println(x, "has been declared")
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	case 7:
		fmt.Println("seven")
	case 8:
		fmt.Println("eight")
	case 9:
		fmt.Println("nine")
	}
}
