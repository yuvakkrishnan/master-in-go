//Leet-Code
package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[rune]int)
	maxLength, start := 0, 0
	for i, char := range s {
		fmt.Println("char", char)
		if lastI, ok := charSet[char]; ok && lastI >= start {
			fmt.Println("lasti", lastI)
			start = lastI
			fmt.Println("lasti after", lastI)

		}
		charSet[char] = i
		maxLength = max(maxLength, i-start+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	value := "aabbbffff"
	result := lengthOfLongestSubstring(value)
	fmt.Println(result)
}
