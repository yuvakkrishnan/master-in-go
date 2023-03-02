// Golang program to return an array
// from a user-defined function

package main

import "fmt"

func GetArray(arr [5]int) [5]int {
	arr[2] = 100
	return arr
}

func main() {
	var intArr [5]int
	var retArr [5]int

	fmt.Println("Enter array elements: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("Element[%d]: ", i)
		fmt.Scanf("%d ", &intArr[i])
	}

	retArr = GetArray(intArr)

	fmt.Println("Array elements: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", retArr[i])
	}
}
