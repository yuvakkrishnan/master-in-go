package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EUR = iota
	GBP
	JPY
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Please provide the amount to be converted")
		return
	}

	rates := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("Invaid Status")
		return
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", amount, rates[EUR]*amount)
	fmt.Printf("%.2f USD is %.2f GBP\n", amount, rates[GBP]*amount)
	fmt.Printf("%.2f USD is %.2f JPY\n", amount, rates[JPY]*amount)
}
