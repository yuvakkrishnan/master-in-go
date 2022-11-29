package main

import "fmt"

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func newParser() parser {
	return parser{sum: make(map[string]result)}

}

func show() {
	fmt.Println(newParser())
}
func main() {
	show()

}
