package main

import "fmt"

type person struct {
	firstName, lastName string
}
type employee struct {
	person
	employeeID int
}

func (p person) details() {
	fmt.Println(p, " "+"My Self")
}

func (e employee) empDetails() {
	fmt.Println(e, " "+"Employee Details !!!")
}

func main() {
	dataPut := person{firstName: "Yuvak", lastName: "Krishnan"}
	dataPut.details()
	empPut := employee{person: person{firstName: "Karishma", lastName: "Krish"}, employeeID: 14252515}
	empPut.empDetails()
}
