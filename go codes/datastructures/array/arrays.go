package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr[0]) // output: 1
	fmt.Println(arr[2]) // output: 3

	for i, val := range arr {
		fmt.Println(i, val)
	}
}
