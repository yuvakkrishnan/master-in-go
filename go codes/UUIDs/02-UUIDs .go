package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// generate a new UUID
	uuid := uuid.NewV4()

	fmt.Println(uuid.String())
}
