package main

import "fmt"

func main() {
	type text struct {
		title string
		word  int
	}

	type book struct {
		text text
		isbn string
	}

	mybook := book{
		text: text{title: "HelloWorld!!"},
		isbn: "Catch up you later",
	}

	fmt.Printf("%s has %d words (isbn :%s)\n", mybook.text.title, mybook.text.word, mybook.isbn)

}
