// Go program to illustrate how to
// create and initialize a map
// Using make() function
package main

import "fmt"

func main() {

	// Creating a map
	// Using make() function
	var My_map = make(map[float64]string)
	fmt.Println(My_map)

	// As we already know that make() function
	// always returns a map which is initialized
	// So, we can add values in it
	My_map[1.3] = "Rohit"
	My_map[1.5] = "Sumit"
	fmt.Println(My_map)
}
