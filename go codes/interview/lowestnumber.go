package main

import "fmt"

func ListofArray(s []int) int {
	fmt.Println(s[1])
<<<<<<< HEAD
	lowest := make(s[0], s)
=======
	lowest := s[0]
>>>>>>> 825db90e4d42816866d191d319ada5f3665140f7
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
