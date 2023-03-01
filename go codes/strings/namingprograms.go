package main

import (
	"fmt"
	"strings"
)

func nameConverter(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "_")
	str = strings.Join([]string{str, ".go"}, "")
	return str
}

func main() {
	strvalue := "Longest Palindromic Substring"
	paraPass := nameConverter(strvalue)
	fmt.Println(paraPass)
}

// ex-5 Buffered or Unbuffered Channel
