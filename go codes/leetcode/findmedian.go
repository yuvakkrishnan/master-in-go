package main

import (
	"fmt"
	"sort"
)

func findMedian(num1 []int, num2 []int) float64 {
	mergedLen := len(num1) + len(num2)
	merged := make([]int, mergedLen)
	copy(merged, num1)
	copy(merged[len(num1):], num2)
	sort.Ints(merged)
	var median float64
	mid := len(merged) / 2
	if len(merged)%2 == 0 {
		median = float64(merged[mid-1]+merged[mid]) / 2.0

	} else {
		median = float64(merged[mid])
	}
	return median

}
func main() {
	num1 := []int{1, 3}
	num2 := []int{2}
	result := findMedian(num1, num2)
	fmt.Println(result)
}

// Question:>
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Constraints:

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106
