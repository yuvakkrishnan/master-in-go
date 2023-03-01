package main

import "fmt"

func longestPalindrome(s []string) []string {
	freq := make(map[string]int)
	duplicates := []string{}

	for _, v := range s {
		freq[v]++
	}
	for str, count := range freq {
		if count > 1 {
			duplicates = append(duplicates, str)
		}
	}
	return duplicates
}

func main() {
	s := []string{"babad", "Hello", "bard", "chatGPT", "chatGPT", "babad", "Hello", "bard", "chatGPT", "chatGPT"}
	result := longestPalindrome(s)
	fmt.Println(result)

}
