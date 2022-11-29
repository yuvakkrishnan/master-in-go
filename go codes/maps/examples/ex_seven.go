// Go program to illustrate the
// modification concept in map

package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a map
	m_a_p := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	fmt.Println("Original map: ", m_a_p)

	// Assigned the map into a new variable
	new_map := m_a_p

	// Perform modification in new_map
	new_map[96] = "Parrot"
	new_map[98] = "Pig"

	// Display after modification
	fmt.Println("New map: ", new_map)
	fmt.Println("\nModification done in old map:\n", m_a_p)
}
