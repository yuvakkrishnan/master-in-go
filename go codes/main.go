package main

import "fmt"

func reverse(s string) string {
	letter := []rune(s)
	for i, j := 0, len(letter)-1; i < j; i, j = i+1, j-1 {
		letter[i], letter[j] = letter[j], letter[i]
	}
	return string(letter)
}

func main() {
	d := "Hello"
	e := reverse(d)
	fmt.Println(e)
}
