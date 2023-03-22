// what is longestPalindromicSubstring ?

// The longest palindromic substring is the longest substring in a given string that is a palindrome.
//A palindrome is a string that is the same forwards and backwards. For example, "racecar" is a palindrome because if you reverse the letters, you get the same word.

// Given a string, the task of finding the longest palindromic substring is to find the longest substring that is a palindrome.
//This is a common problem in computer science and has many applications in areas such as natural language processing, data compression, and bioinformatics.

// For example, in the string "babad", the longest palindromic substring is "bab", which is a palindrome.
//Note that "aba" is also a palindrome substring, but it is not as long as "bab".

package main

import "fmt"

func longestPalindromicSubstring(s string) string {
	var maxStart, maxEnd int
	for i := 0; i < len(s); i++ {
		// check for odd-length palindromes
		start, end := expandAroundCenter(s, i, i)
		if end-start > maxEnd-maxStart {
			maxStart, maxEnd = start, end
		}

		// check for even-length palindromes
		start, end = expandAroundCenter(s, i, i+1)
		if end-start > maxEnd-maxStart {
			maxStart, maxEnd = start, end
		}
	}

	return s[maxStart : maxEnd+1]
}

func expandAroundCenter(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func main() {
	s := "cbbd"
	res := longestPalindromicSubstring(s)
	fmt.Println(res) // Output: "bab"
}
