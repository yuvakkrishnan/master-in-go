package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	items := os.Args[1:]
	if len(items) == 0 {
		fmt.Println("Send me some items and I will sort them")
	}
	sort.Strings(items)
	var data []byte
	for i, fck := range items {
		data = strconv.AppendInt(data, int64(i+1), 10)
		data = append(data, '.', ' ')
		data = append(data, fck...)
		data = append(data, '\n')

	}

	err := ioutil.WriteFile("sexy3.txt", data, 0544)
	if err != nil {
		fmt.Println(err)
	}
}
