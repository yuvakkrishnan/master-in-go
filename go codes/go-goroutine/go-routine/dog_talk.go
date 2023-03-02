package main

import "fmt"

type Animal interface {
	speak() string
}

type Dog struct{}

func (d Dog) speak() string {
	return "woof"
}

type Cat struct{}

func (c Cat) speak() string {
	return "Meow!!"
}

func talk(a Animal, message chan<- string) {
	message <- a.speak()
}

func main() {
	messages := make(chan string)

	// for i := 1; i <= 2; i++ {
	// 	go talk(Dog{}, messages)
	// }
	go talk(Dog{}, messages)
	go talk(Cat{}, messages)

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
