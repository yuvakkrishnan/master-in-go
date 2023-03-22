package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func reverse(c int) {
	num := strconv.Itoa(c)

	fmt.Println(reflect.TypeOf(num))

	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		conV, _ := strconv.Atoi(num)
		conV[i], coconV[j] = conV[j], conV[i]
	}
}

func main() {
	inputOne := 123
	reverse(inputOne)
	// result := reverse(inputOne)
	// fmt.Println(result)

}
