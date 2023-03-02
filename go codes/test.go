package main

import "fmt"

func duplicates(num []int) []int {
	count := make(map[int]int)

	for _, v := range num {
		count[v]++
	}

	var duplicates []int

	for t, y := range count {
		if y > 1 {
			duplicates = append(duplicates, t)
		}
	}
	return duplicates
}

func main() {

	numBers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	dup := duplicates(numBers)
	fmt.Println(dup)
}
