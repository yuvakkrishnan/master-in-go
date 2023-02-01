/*
1. Two Sum
Source: https://leetcode.com/problems/two-sum/
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import "fmt"

// Time complexity: O(n)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	//HAsh map
	hm := make(map[int]int)

	for i, v := range nums {
		_, ok := hm[v]
		if ok {
			return []int{i, hm[v]}
		}
		hm[target-v] = i

	}
	return nil
}

func main() {
	intArr := []int{1, 2, 3, 4, 5, 6}
	targetNum := 100

	result := twoSum(intArr, targetNum)

	fmt.Printf("%d Final Result Value", result)
}

// Second solution
// brute force
// Time complexity: O(n^2)
// Space complexity: O(1)
// func twoSum1(nums []int, target int) []int {
// 	for i, j := range nums {
// 		for k := i + 1; k < len(nums); k++ {
// 			if nums[k]+j == target {
// 				return []int{i, k}
// 			}
// 		}
// 	}
// 	return []int{}
// }
