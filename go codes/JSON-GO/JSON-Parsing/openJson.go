package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileJson, err := os.Open("user-test.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully file opened!!!!")

	bytesValue, _ := ioutil.ReadAll(fileJson)

	fmt.Println(bytesValue)
}
