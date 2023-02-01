package main

import (
	"fmt"
	"strings"
)

func reverseOrder(str string) string {

	l := strings.Split(str, "")
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]

	}
	return strings.Join(l, " ")
}
func main() {

	g := reverseOrder("Hekko")
	fmt.Println(g)
}
