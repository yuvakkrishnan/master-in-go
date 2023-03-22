package main

// import (
// 	"fmt"
// 	"os"
// )

// var NumberToWord = map[int]string{
// 	1:  "one",
// 	2:  "two",
// 	3:  "three",
// 	4:  "four",
// 	5:  "five",
// 	6:  "six",
// 	7:  "seven",
// 	8:  "eight",
// 	9:  "nine",
// 	10: "ten",
// }

// func convert1to10(n int) (w string) {
// 	if n < 10 {
// 		w = NumberToWord[n]
// 		return
// 	}

// 	r := n % 10
// 	if r == 0 {
// 		w = NumberToWord[n]
// 	} else {
// 		w = NumberToWord[n-r] + "-" + NumberToWord[r]
// 	}
// 	return
// }

// func main() {

// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		fmt.Println("Type Number 1-10")
// 		return
// 	}
// 	if args == NumberToWord{
// 		fmt.Println("a")
// 	}
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(convert1to10(i))
// 	}
// }
