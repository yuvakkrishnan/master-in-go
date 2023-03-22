package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {

	args := os.Args[1:]
	if len(args) != 0 {
		fmt.Println("Send me some items and I will sort them")
	}
	sort.Strings(args)
	var sortFile []byte

	for _, data := range args {
		sortFile = append(sortFile, data...)
		sortFile = append(sortFile, '\n')
	}
	// fmt.Println(sortFile) O\P [97 112 112 108 101 10 98 97 110 97 110 97 10 103 114 97 112 114 101 10 106 117 105 99 101 10 111 114 97 110 103 101 10]

	err := ioutil.WriteFile("sorted.txt", sortFile, 0642)
	if err != nil {

		fmt.Println(err)
		return
	}

}
