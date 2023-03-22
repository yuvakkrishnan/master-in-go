// Go program to illustrate how to add
// a key-value pair in the map using
// make() function
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

	// Adding new key-value pairs in the map
	m_a_p[95] = "Parrot"
	m_a_p[96] = "Crow"
	fmt.Println("Map after adding new key-value pair:\n", m_a_p)

	// Updating values of the map
	m_a_p[91] = "PIG"
	m_a_p[93] = "DONKEY"
	fmt.Println("\nMap after updating values of the map:\n", m_a_p)
}
