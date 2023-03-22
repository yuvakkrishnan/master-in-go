// Golang program to demonstrate the
// example of make() function

package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	fmt.Printf("a: Type: %T, Length: %d, Capacity: %d\n",
		a, len(a), cap(a))
	fmt.Println("value of a:", a)

	b := make([]int, 10, 20)
	fmt.Printf("b: Type: %T, Length: %d, Capacity: %d\n",
		b, len(b), cap(b))
	fmt.Println("value of b:", b)

	c := make([]int, 0, 5)
	fmt.Printf("c: Type: %T, Length: %d, Capacity: %d\n",
		c, len(c), cap(c))
	fmt.Println("value of c:", c)
}
