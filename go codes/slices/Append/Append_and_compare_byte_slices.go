package main

import (
	"bytes"
	"fmt"
)

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}
	header = append(header, 'F', 'G', 'Y')

	if bytes.Equal(png, header) {
		fmt.Print("Both are Equal")
	} else {
		fmt.Println("Botha are not Equal")
	}
}
