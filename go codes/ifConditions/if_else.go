package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the Grade : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade > 100 {
		status = "Grade is out of range "
	} else if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"

	}
	fmt.Println("A grade of", grade, "is", status)

}
