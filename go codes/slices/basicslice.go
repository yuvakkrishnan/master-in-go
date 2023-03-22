package main

import "fmt"

func main() {

	var books [5]string
	books[0] = "dragulla"
	books[1] = "HelloWorld"
	books[2] = "Mudamada"

	games := []string{"Cricket", "Football"}

	fmt.Printf("Books :%#v \n", books)
	fmt.Printf("Games :%#v \n", games)
	fmt.Printf("game  :%T\n", games)
	fmt.Printf("games :len :%d\n", len(games))
	fmt.Printf("nil?  : %t\n", games == nil)

}
