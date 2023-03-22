// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const corpus = "" + "lazy cat jump again and again and again"

// func main() {
// 	words := strings.Fields(corpus)
// 	query := os.Args[1:]
// 	// quries: 							Label break from parent loop
// 	for _, f := range query {
// 		fmt.Printf("%q\n", f)

// 		for s, p := range words {
// 			// fmt.Printf("%q %q", s, p)
// 			if f == p {
// 				fmt.Printf("# %-2d: %q\n",
// 					s+1, p)
// 				break quries

// 			}
// 		}
// 	}

}

// ******************************************************   Break from a Switch using Labels	***************************************************

// ******************************************************goto jump statement***************************************************

// import (
// 	"fmt"
// )

// func main() {
// 	var i int
// loop:
// 	if i < 3 {
// 		fmt.Printf(" # % d <--> Hello", i)

// 		i++
// 		goto loop
// 	}

// 	fmt.Println("Done")
// }
// ******************************************************goto jump statement***************************************************
