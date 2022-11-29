// Golang Program to illustrate the usage of
// range keyword over maps in Golang

package main

import "fmt"

// main function
func main() {

	// Creating map of student ranks
	student_rank_map := map[string]int{
		"Nidhi": 3,
		"Nisha": 2,
		"Rohit": 1,
	}

	// Printing map using keys only
	for student := range student_rank_map {
		fmt.Println("Rank of", student, "is: ",
			student_rank_map[student])
	}

	// Printing maps using key-value pair
	for student, rank := range student_rank_map {
		fmt.Println("Rank of", student, "is: ", rank)
	}
}
