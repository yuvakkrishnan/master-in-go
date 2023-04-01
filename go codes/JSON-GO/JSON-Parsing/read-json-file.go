package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users`
}

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}
type Geo struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	Catchphrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

func main() {

	fileOpen, err := os.Open("users.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Sucessfully file opened users.json!!!")

	byteValue, _ := ioutil.ReadAll(fileOpen)
	var users Users
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("user name :" + users.Users[i].Name)
	}
}
