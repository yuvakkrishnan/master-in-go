package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Print("Enter Word To Execute !!!!")
		return
	}
	query := args[0]
	english := []string{"one", "Two", "Three"}
	tamil := []string{"ondru", "rendu", "mondru"}

	for i, lang := range english {
		if query == lang {

			fmt.Printf("%q mean %q\n", lang, tamil[i])
			return
		}
	}
	fmt.Printf("%q Not Found Query\n", query)
}
