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
	strvalue := "example-2 multiple Goroutines "
	paraPass := nameConverter(strvalue)
	fmt.Println(paraPass)
}
