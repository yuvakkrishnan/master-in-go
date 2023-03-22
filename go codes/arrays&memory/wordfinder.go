package main

import (
	"fmt"
	"os"
	"strings"
)

const Corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {

	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Println("Enter the word to find the word")
		return
	}

	filter := [...]string{
		"and", "or", "was", "the", "since", "very",
	}

	word := strings.Fields(Corpus)

quries:
	for _, q := range query {
		q := strings.ToLower(q)
		for _, v := range filter {
			if v == q {
				continue quries
			}
		}

		for i, w := range word {
			if q == w {
				fmt.Printf("#%-2d %q\n", i+1, w)
				break
			}
		}
	}

}
