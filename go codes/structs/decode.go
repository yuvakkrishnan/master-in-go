package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	name       string
	permission map[string]bool
}

func main() {
	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user
	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println("Error Occured !!!")
		return
	}
}
