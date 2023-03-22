package main

import "fmt"

var N = 3

func main() {

	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		m[i] = &i
	}

	fmt.Println("Value1: ", m[0])

	//Value1:012

	for _, v := range m {
		fmt.Println("Value2: ", *v)
		//Value2:)012
	}
}
