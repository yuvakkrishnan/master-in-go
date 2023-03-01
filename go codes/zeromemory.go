package main

import (
	"fmt"
)

func longestPalindromicSubstring(s string) {

	var maxStart, maxEnd int
	for i := 0; i < len(s); i++ {
		// fmt.Println("s", s, "i", i, "i2", i)
		start, end := expandAround(s, i, i)
		fmt.Println()
		if end-start > maxEnd-maxStart {
			maxStart, maxEnd = start, end
		}
	}

	// fmt.Println(s[maxStart : maxEnd+1])

}

func expandAround(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {

		left--
		right++
	}

	return left + 1, right - 1
}

func main() {
	str := "HelloWorld"
	longestPalindromicSubstring(str)
	//result :=
	//fmt.Println(result)

}
