package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "H#ll* @orl%"
	fixed := strings.ReplaceAll(broken, "#", "e")
	fixed = strings.ReplaceAll(fixed, "*", "o")
	fixed = strings.ReplaceAll(fixed, "@", "W")
	fixed = strings.ReplaceAll(fixed, "%", "d")

	fmt.Println(fixed) // Output: Hello world%
}
