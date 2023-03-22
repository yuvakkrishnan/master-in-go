package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "\t Hello, World\n "
	fmt.Printf("%d %q\n", len(s), s)
	t := strings.TrimSpace(s)
	fmt.Printf("%d %q\n", len(t), t)
}
