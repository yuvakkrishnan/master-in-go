package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if a := os.Args; len(a) != 2 {
		fmt.Printf("Give me a number ")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot Convert %q \n", a[1])
	} else {
		fmt.Printf("%s *2  %d \n", a[1], n*2)
	}

}
