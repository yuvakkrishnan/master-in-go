// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	i := 2
// 	fmt.Print("Write ", i, " as ")
// 	switch i {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	}

// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It's the weekend")
// 	default:
// 		fmt.Println("It's a weekday")
// 	}

// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("It's before noon")
// 	default:
// 		fmt.Println("It's after noon")
// 	}

// 	whatAmI := func(i interface{}) {
// 		switch t := i.(type) {
// 		case bool:
// 			fmt.Println("I'm a bool")
// 		case int:
// 			fmt.Println("I'm an int")
// 		default:
// 			fmt.Printf("Don't know type %T\n", t)
// 		}
// 	}
// 	whatAmI(true)
// 	whatAmI(1)
// 	whatAmI("hey")
// }
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const usage = `[command] [string]
// 	Available commands: lower, upper and title`

// func main() {
// 	args := os.Args
// 	if len(os.Args) != 3 {
// 		fmt.Println(usage)
// 		return
// 	}
// 	cmd, str := args[1], args[2]
// 	switch cmd {
// 	case "lower":
// 		fmt.Printf(strings.ToLower(str))
// 	case "upper":
// 		fmt.Println(strings.ToUpper(str))
// 	case "titleName":
// 		fmt.Println(strings.ToTitle(str))

// 	}

// }
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {

// 	if len(os.Args) != 2 {
// 		fmt.Println("give me a month name")
// 		return
// 	}
// 	year := time.Now().Year()
// 	leap := year%4 == 0 && (year%100 == 0 || year%400 == 0)

// 	days, month := 28, os.Args[1]

// 	switch strings.ToLower(month) {
// 	case "apr", "june", "sep", "nov":
// 		days = 30
// 	case "jan", "march", "may", "july":
// 		days = 31
// 	case "feb":
// 		if leap {
// 			days = 29
// 		}
// 	default:
// 		fmt.Printf("%q is not a month.\n", month)
// 		return

// 	}
// 	fmt.Printf("%q has %d days.\n", month, days)

// }

package main
