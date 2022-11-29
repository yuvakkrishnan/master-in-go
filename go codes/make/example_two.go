// Golang program to demonstrate the
// example of make() function

package main

import (
	"fmt"
)

func main() {
	// Creating a map using make()
	var student = make(map[string]int)

	// Assigning
	student["Alvin"] = 21
	student["Alex"] = 47
	student["Mark"] = 27

	// Printing the map, its type and lenght
	fmt.Println(student)
	fmt.Printf("Type: %T, Length: %d\n",
		student, len(student))
}
