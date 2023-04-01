package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	// generate a new random UUID
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(err)
	}

	// set the version and variant bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // variant 1

	// print the UUID as a string
	fmt.Printf("%x-%x-%x-%x-%x\n", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
