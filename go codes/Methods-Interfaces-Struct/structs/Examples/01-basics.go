package main

import "fmt"

type person struct {
	name, lastNAme string
	age            int
}

func main() {

	entryOne := person{name: "Sarvan", lastNAme: "KokiKumar", age: 21}
	s := entryOne.age
	fmt.Println(s)
	fmt.Printf("%+v ", entryOne)
}
