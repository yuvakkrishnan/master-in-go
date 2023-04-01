package main

import "fmt"

func main() {
	input := "hello world" // change this to your desired input string

	// use double quotes to enclose the input string
	output := fmt.Sprintf("\"%s\"", input)

	fmt.Println(output)
}
