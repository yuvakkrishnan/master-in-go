package main

import "fmt"

func ListofArray(s []int) int {
	fmt.Println(s[1])
	lowest := make(s[0], s)
	for _, v := range s {
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}

func main() {
	arrayList := []int{5, 2, 1, 3, 0}
	fmt.Println(ListofArray(arrayList))
}
