// Go program to illustrate how to
// create and initialize maps
package main

import "fmt"

func main() {

	// Creating and initializing empty map
	// Using var keyword
	var map_1 map[int]int

	// Checking if the map is nil or not
	if map_1 == nil {

		fmt.Println("True")
	} else {

		fmt.Println("False")
	}

	// Creating and initializing a map
	// Using shorthand declaration and
	// using map literals
	map_2 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
<<<<<<< HEAD
	for key, value := range map_2 {
		fmt.Println("Key:", key, "Value:", value)

	}
	fmt.Println("Map-2: ", map_2)

=======

	fmt.Println("Map-2: ", map_2)
>>>>>>> 825db90e4d42816866d191d319ada5f3665140f7
}
