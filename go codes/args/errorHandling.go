// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {

// 	// It's never bring error

// 	s := strconv.Itoa(42)
// 	fmt.Println(s)
// }

// Atoi sometimes fails
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	s, err := strconv.Atoi(os.Args[1])

// 	fmt.Println("Geting numbers", s)
// 	fmt.Println("Erro value", err)
// }

// Handling Error

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )
// func main {
// 	age := os.Args[1]

// 	n,err:=strconv.Atoi(age)
// 	if err != nil {
// 		fmt.Println("Error",err)
// 		return
// 	}
// 	fmt.Printf("sucess :Converted %q to %d ./n",age,n)
// }
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// const usage = `Feet to meters
// 	______________
// 	This Feet to meter program
// 	usage :
// 	feet[feetsToConvert]`

// func main() {

// 	if len(os.Args) < 2 {
// 		fmt.Println(strings.TrimSpace(usage))
// 		return
// 	}
// 	args := os.Args[1]

// 	feet, _ := strconv.ParseFloat(args, 64)
// 	//ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:

// 	meter := feet * 0.3848

// 	fmt.Printf("%g feet is %g meter.\n", feet, meter)
// }
// *********************************Movie ratings

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if a := os.Args; len(a) != 2 {
// 		fmt.Println("Requires age")
// 		return

// 	} else if n, err := strconv.Atoi(a[1]); err != nil || n < 0 {
// 		fmt.Printf(`Other characters Cannot be age: %q`+"\n ", os.Args[1])
// 		return

// 	} else if n >= 13 && n <= 17 {
// 		fmt.Printf("PG-13")
// 	} else if n > 17 {
// 		fmt.Printf("R-Rated")
// 	} else if n < 13 {
// 		fmt.Printf("PG-Rated")
// 	}
// }
// *********************************Movie ratings

// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if eod := os.Args; len(eod) < 2 {
// 		fmt.Println("Pick a number")
// 		return
// 	} else if engal, err := strconv.Atoi(eod[1]); err != nil {
// 		fmt.Printf("%d is not a number", engal)
// 	} else if engal%2 == 0 {
// 		fmt.Printf("%d is even", engal)

// 	} else {
// 		fmt.Printf("%d is odd", engal)
// 	}

// }
// **************************************************************************************

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Give me a month name")
// 		return
// 	}

// 	// get the current year and find out whether it's a leap year
// 	year := time.Now().Year()
// 	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

// 	// setting it to 28, saves me typing it below again
// 	days := 28

// 	month := os.Args[1]

// 	// case insensitive
// 	if m := strings.ToLower(month); m == "april" ||
// 		m == "june" ||
// 		m == "september" ||
// 		m == "november" {
// 		days = 30
// 	} else if m == "january" ||
// 		m == "march" ||
// 		m == "may" ||
// 		m == "july" ||
// 		m == "august" ||
// 		m == "october" ||
// 		m == "december" {
// 		days = 31
// 	} else if m == "february" {
// 		// try a "logical and operator" above.
// 		// like: `else if m == "february" && leap`
// 		if leap {
// 			days = 29
// 		}
// 	} else {
// 		fmt.Printf("%q is not a month.\n", month)
// 		return
// 	}

// 	fmt.Printf("%q has %d days.\n", month, days)
// }

package main
