package main

import "fmt"

func changeX(x *int) {
	*x = 0
}
func main() {
	var x = 200

	changeX(&x)

	fmt.Println("After func Calls prints", x)

}
