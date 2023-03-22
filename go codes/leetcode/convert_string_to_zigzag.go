package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([]string, numRows)

	index, step := 0, 1
	for _, v := range s {
		rows[index] += string(v)
		if index == 0 {
			step = 1
		} else if index == numRows-1 {
			step = -1
		}

		index += step
	}
	var result string
	for _, v := range rows {
		result += v
	}

	return result
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 7
	result := convert(s, numRows)
	fmt.Println(result)
}
