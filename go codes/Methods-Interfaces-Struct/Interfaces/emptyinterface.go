package main

import "fmt"

func printType(i interface{}) {
	fmt.Printf("Type :%T ,value :%v\n", i, i)
}

func main() {
	printType("Helloworld")
	printType(123)
	printType(true)
	printType(struct{}{})

}
