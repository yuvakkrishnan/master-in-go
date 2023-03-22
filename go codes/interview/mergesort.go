package main

import "fmt"

func twoSort(a, b []int) []int {

	result := make([]int, len(a)+len(b))
	i := 0
	j := 0
	k := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[i] {
			result[k] = a[i]
			i++
		} else {
			result[k] = b[j]
			j++
		}
		k++

	}

	for i < len(a) {
		result[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		result[k] = b[j]
		j++
		k++
	}
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	a := mergeSort(arr[:mid])
	b := mergeSort(arr[mid:])
	return twoSort(a, b)
}
func main() {
	arr := []int{5, 2, 7, 3, 9, 1, 8, 6, 4}
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)
}
