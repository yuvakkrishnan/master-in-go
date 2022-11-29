package main

import (
	"fmt"
)

func main() {
	nums := []int{9, 7, 5}
	nums = append(nums, 2, 4, 6)

	fmt.Println(nums[len(nums)-2:])
}
