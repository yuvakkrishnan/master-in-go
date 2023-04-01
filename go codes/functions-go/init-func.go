package main

import "fmt"

const count int = 10

func init() {

	for i := 0; i < count; i++ {
		fmt.Println("Workking First", i)
	}
}

func main() {
	fmt.Println("Function Ends!!!!")
}
