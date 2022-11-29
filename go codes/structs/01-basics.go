package main

import "fmt"

func main() {
	type Person struct {
		name, lastNAme string
		age            int
	}
	age := Person.age

	entryOne := Person{name: "Sarvan", lastNAme: "KokiKumar", age: 21}

	fmt.Printf("%+v ", entryOne)
}
