// Go program to illustrate how to
// retrieve the value of the key

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

	// Retrieving values with the help of keys
	value_1 := m_a_p[90]
	value_2 := m_a_p[93]
	fmt.Println("Value of key[90]: ", value_1)
	fmt.Println("Value of key[93]: ", value_2)
}
