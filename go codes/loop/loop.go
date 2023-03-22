package main

// import (
// 	"fmt"
// )
// func main() {
// 	var sum int
// 	// sum += 1
// 	// sum += 2
// 	// sum += 3
// 	// sum += 4
// 	// sum += 5

// 	for i := 1; i <= 1000; i++ {
// 		sum += i
// 	}

// 	fmt.Println(sum)
// }

// *************************************************************************************//Basic Loop and  Break Statement// ************************************************************

// BAasic Loop structure
// package main

// import "fmt"

// func main() {
// 	var (
// 		sum int
// 		i   = 1
// 	)

// 	for {
// 		if i > 5 {
// 			break
// 		}
// 		sum = sum + i
// 		fmt.Println(i, "-->", sum)
// 		i++
// 	}
// 	fmt.Println(sum)
// }																										   output
//    i   sum
// 1 --> 1
// 2 --> 3
// 3 --> 6
// 4 --> 10
// 5 --> 15
// 15

// *************************************************************************************//Basic Loop and  Break Statement// ************************************************************

// *************************************************************************************//Nested Loop // ************************************************************

// package main

// import "fmt"

// const max = 5

// func main() {
// 	fmt.Printf("%5s", "X") act a as x header table string
// 	for i := 0; i <= max; i++ {   ***It will print 0,1,2,3,4,5 in Horizontally
// 		fmt.Printf("%5d", i)

// 	}
// 	fmt.Println()
// 	for i := 0; i <= max; i++ {			 ***It will print 0,1,2,3,4,5 in Verically
// 		fmt.Printf("%5d", i)
// 		for j := 0; j <= max; j++ {
// 			fmt.Printf("%5d", i*j)			****Mulitiply the inner values
// 		}
// 		fmt.Println()
// 	}
// }

// Nested Loop

// X    0    1    2    3    4    5
// 0    0    0    0    0    0    0
// 1    0    1    2    3    4    5
// 2    0    2    4    6    8   10
// 3    0    3    6    9   12   15
// 4    0    4    8   12   16   20
// 5    0    5   10   15   20   25

// *************************************************************************************//Nested Loop // ************************************************************
// *************************************************************************************// Loop Over slice// ************************************************************

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// const max = 10

// func main() {
// 	word := strings.Fields("lazy cat jumps again and again ")

// 	for h := 0; h < len(word); h++ {
// 		fmt.Printf("# %2d and %q\n", h+1, word[h])
// 	}

// }
// *************************************************************************************// Loop Over slice// ************************************************************

// *************************************************************************************// For Range Loop  slice// ************************************************************

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// for i, v := range os.Args {
// 	if i == 0 {
// 		continue
// 	}
// 	fmt.Printf("%q\n", v)
// }
// 	for _, v := range os.Args[1:4] {

// 		fmt.Printf("%q\n", v)
// 	}
// }
// *************************************************************************************// For Range Loop  slice// ************************************************************

// *************************************************************************************// For Loop slice Range EXE// ************************************************************
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {

// 	args := os.Args
// 	if len(os.Args) != 2 {
// 		fmt.Println(" Give me the size of the table")
// 		return
// 	}
// 	size, err := strconv.Atoi(args[1])
// 	if err != nil || size < 0 {
// 		fmt.Println("Worng Size")
// 	}
// 	fmt.Printf("%5s ", "x")
// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)
// 	}
// 	fmt.Println()
// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)
// 		for j := 0; j <= size; j++ {
// 			fmt.Printf("%5d", i*j)
// 		}
// 		fmt.Println()

// 	}
// }

// *************************************************************************************// For Loop slice Range EXE// ************************************************************
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// const (
// 	validOps = "%*/+-"

// 	usageMsg       = "Usage: [op=" + validOps + "] [size]"
// 	sizeMissingMsg = "Size is missing"
// 	invalidOpMsg   = `Invalid operator.
// Valid ops one of: ` + validOps

// 	invalidOp = -1
// )

// func main() {
// 	args := os.Args[1:]

// 	switch l := len(args); {
// 	case l == 1:
// 		fmt.Println(sizeMissingMsg)
// 		fallthrough
// 	case l < 1:
// 		fmt.Println(usageMsg)
// 		return
// 	}

// 	op := args[0]
// 	if strings.IndexAny(op, validOps) == invalidOp {
// 		fmt.Println(invalidOpMsg)
// 		return
// 	}

// 	size, err := strconv.Atoi(args[1])
// 	if err != nil || size <= 0 {
// 		fmt.Println("Wrong size")
// 		return
// 	}

// 	fmt.Printf("%5s", op)
// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)
// 	}
// 	fmt.Println()

// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)

// 		for j := 0; j <= size; j++ {
// 			var res int

// 			switch op {
// 			case "*":
// 				res = i * j
// 			case "%":
// 				if j != 0 {
// 					res = i % j
// 				}
// 			case "/":
// 				if j != 0 {
// 					res = i / j
// 				}
// 			case "+":
// 				res = i + j
// 			case "-":
// 				res = i - j
// 			}

// 			fmt.Printf("%5d", res)
// 		}
// 		fmt.Println()
// 	}
// }
// ************************************************************************************************
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// const (
// 	validOps = "%*/+-"

// 	usageMsg       = "Usage: [op=" + validOps + "] [size]"
// 	sizeMissingMsg = "Size is missing"
// 	invalidOpMsg   = `Invalid operator.
// Valid ops one of: ` + validOps

// 	invalidOp = -1
// )

// func main() {
// 	args := os.Args[1:]

// 	switch l := len(args); {
// 	case l == 1:
// 		fmt.Println(sizeMissingMsg)
// 		fallthrough
// 	case l < 1:
// 		fmt.Println(usageMsg)
// 		return
// 	}

// 	op := args[0]
// 	if strings.IndexAny(op, validOps) == invalidOp {
// 		fmt.Println(invalidOpMsg)
// 		return
// 	}

// 	size, err := strconv.Atoi(args[1])
// 	if err != nil || size <= 0 {
// 		fmt.Println("Wrong size")
// 		return
// 	}

// 	fmt.Printf("%5s", op)
// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)
// 	}
// 	fmt.Println()

// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)

// 		for j := 0; j <= size; j++ {
// 			var res int

// 			switch op {
// 			case "*":
// 				res = i * j
// 			case "%":
// 				if j != 0 {
// 					res = i % j
// 				}
// 			case "/":
// 				if j != 0 {
// 					res = i / j
// 				}
// 			case "+":
// 				res = i + j
// 			case "-":
// 				res = i - j
// 			}

// 			fmt.Printf("%5d", res)
// 		}
// 		fmt.Println()
// 	}
// }
// ##################################################02-sum-the-numbers-verbose################################################

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	const min, max = 1, 10
// 	var sum int

// 	for i := min; i <= max; i++ {
// 		sum += i
// 		fmt.Print(i)
// 		// if i < max {
// 		// 	fmt.Print("+")
// 		// }

// 	}
// 	// fmt.Printf("= %d\n", sum)
// }
// ##################################################02-sum-the-numbers-verbose################################################

// ####################################03-sum-up-to-n$+######################

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("gimme two numbers")
// 		return
// 	}

// 	min, err1 := strconv.Atoi(os.Args[1])
// 	max, err2 := strconv.Atoi(os.Args[2])
// 	if err1 != nil || err2 != nil || min >= max {
// 		fmt.Println("wrong numbers")
// 		return
// 	}

// 	var sum int
// 	for i := min; i <= max; i++ {
// 		sum += i

// 		fmt.Print(i)
// 		if i < max {
// 			fmt.Print(" + ")
// 		}
// 	}
// 	fmt.Printf(" = %d\n", sum)
// }
// ####################################03-sum-up-to-n$+######################

// ####################################Continue Loop######################

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("gimme two numbers")
// 		return
// 	}

// 	min, err1 := strconv.Atoi(os.Args[1])
// 	max, err2 := strconv.Atoi(os.Args[2])
// 	if err1 != nil || err2 != nil || min >= max {
// 		fmt.Println("wrong numbers")
// 		return
// 	}

// 	var sum int
// 	for i := min; i <= max; i++ {
// 		if i%2 != 0 {
// 			continue
// 		}
// 		sum += i

// 		fmt.Print(i)
// 		if i < max-1 {
// 			fmt.Print(" + ")
// 		}
// 	}
// 	fmt.Printf(" = %d\n", sum)
// }

// ####################################Continue Loop######################

// ####################################infinite-kill######################
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	for {
// 		var c string

// 		switch rand.Intn(4) {
// 		case 0:
// 			c = "\\"
// 		case 1:
// 			c = "/"
// 		case 2:
// 			c = "|"
// 		case 3:
// 			c = "-"
// 		}
// 		fmt.Printf("\r%s Please Wait. Processing....", c)
// 		time.Sleep(time.Millisecond * 250)
// 	}
// }
// ####################################infinite-kill######################

// ####################################Break_Loop######################

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("gimme two numbers")
// 		return
// 	}

// 	min, err1 := strconv.Atoi(os.Args[1])
// 	max, err2 := strconv.Atoi(os.Args[2])
// 	if err1 != nil || err2 != nil || min >= max {
// 		fmt.Println("wrong numbers")
// 		return
// 	}

// 	var (
// 		i   = min
// 		sum int
// 	)

// 	for {
// 		if i > max {
// 			break
// 		} else if i%2 != 0 {
// 			i++
// 			continue
// 		}

// 		fmt.Print(i)
// 		if i < max-1 {
// 			fmt.Print(" + ")
// 		}

// 		sum += i
// 		i++
// 	}
// 	fmt.Printf(" = %d\n", sum)
// }
// ####################################Break_Loop######################
