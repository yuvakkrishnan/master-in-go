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
<<<<<<< HEAD
	strvalue := "Convert string to zigzag"
=======
	strvalue := "Longest Palindromic Substring"
>>>>>>> 825db90e4d42816866d191d319ada5f3665140f7
	paraPass := nameConverter(strvalue)
	fmt.Println(paraPass)
}

// ex-5 Buffered or Unbuffered Channel
