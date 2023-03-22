package main

type Salary struct {
	Basic, HRA, TA float64
}
type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MOnthlysalary              []Salary
}
