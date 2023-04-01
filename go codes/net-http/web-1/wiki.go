package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	fileOpen, _ := os.Open("lorem.txt")
	readFile, _ := ioutil.ReadAll(fileOpen)

	p1 := &Page{Title: "Test", Body: []byte(string(readFile))}
	p1.save()
	p2, _ := loadPage("Test")
	fmt.Println(string(p2.Body))
}
