package main

import "fmt"

func twoSum(nums []int, target int) []int {
	final := make(map[int]int)
	for k, o := range nums {
		reduce := target - o
		if tux, ok := final[reduce]; ok {
			return []int{tux, k}

		}
		final[o] = k

	}
	return []int{}
}

func main() {
	intArr := []int{4, 5, 6, 7, 8}
	targetNum := 9

	result := twoSum(intArr, targetNum)
	fmt.Printf("%d Result", result)
}
