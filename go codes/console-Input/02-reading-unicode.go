package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}
	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A Pressed!!!")
	case 'a':
		fmt.Print("a pressed")
		break
	}
}
