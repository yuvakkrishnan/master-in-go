package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Users struct {
	Users []User `json:"users"`
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
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Company struct {
	Name        string `json:"name"`
	Catchphrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

func jsonFileOpen(filename string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully File Opened users.json !!!")
	defer jsonFile.Close()
	var data map[string]interface{}
	err = json.NewDecoder(jsonFile).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func jsonFileRead(data map[string]interface{}) {
	usersData := data["users"].([]interface{})
	users := make([]User, len(usersData))
	for i, userData := range usersData {
		user := userData.(map[string]interface{})
		users[i].Id = int(user["id"].(float64))
		users[i].Name = user["name"].(string)
		users[i].Username = user["username"].(string)
		users[i].Email = user["email"].(string)

		addressData := user["address"].(map[string]interface{})
		users[i].Address.Street = addressData["street"].(string)
		users[i].Address.Suite = addressData["suite"].(string)
		users[i].Address.City = addressData["city"].(string)
		users[i].Address.Zipcode = addressData["zipcode"].(string)

		geoData := addressData["geo"].(map[string]interface{})
		users[i].Address.Geo.Lat = geoData["lat"].(string)

		users[i].Address.Geo.Lng = geoData["lng"].(string)

		companyData := user["company"].(map[string]interface{})
		users[i].Company.Name = companyData["name"].(string)
		users[i].Company.Catchphrase = companyData["catchPhrase"].(string)
		users[i].Company.Bs = companyData["bs"].(string)
	}
	fmt.Println(users)
}

func main() {
	data, err := jsonFileOpen("users.json")
	if err != nil {
		panic(err)
	}
	jsonFileRead(data)
}
