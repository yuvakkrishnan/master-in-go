package main

// ###################################### Random & Seed initiating ######################################

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	rand.Seed(10)
// 	guess := 10

// 	for i := 0; i != guess; {
// 		i = rand.Intn(guess + 1)
// 		fmt.Printf("%d", i)
// 	}

// }
// ###################################### Random & Seed ######################################
// ###################################### Random & Seed & Time Unix  initiating ######################################

// import (
// 	"fmt"

// 	"math/rand"

// 	"time"
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	guess := 10

// 	for n := 0; n != guess; {
// 		// n = rand.Intn(guess + 1)
// 		// fmt.Printf("%d", n)
// 	}
// 	fmt.Println()
// }
// ###################################### Random & Seed & Time Unix  initiating ######################################
// ##########################################################	First Turn Winner	############################################################################
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"
// )

// const (
// 	maxTurns = 5 // less is more difficult
// 	usage    = `Welcome to the Lucky Number Game! 🍀
// The program will pick %d random numbers.
// Your mission is to guess one of those numbers.
// The greater your number is, harder it gets.
// Wanna play?
// `
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	args := os.Args[1:]

// 	if len(args) != 1 {
// 		fmt.Printf(usage, maxTurns)
// 		return
// 	}

// 	guess, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		fmt.Println("Not a number.")
// 		return
// 	}

// 	if guess <= 0 {
// 		fmt.Println("Please pick a positive number.")
// 		return
// 	}

// 	for turn := 1; turn <= maxTurns; turn++ {
// 		n := rand.Intn(guess) + 1

// 		// Better, why?
// 		//
// 		// Instead of nesting the if statement into
// 		//   another if statement; it simply continues.
// 		//
// 		// TLDR: Avoid nested statements.
// 		if n != guess {
// 			continue
// 		}

// 		if turn == 1 {
// 			fmt.Println("🥇 FIRST TIME WINNER!!!")
// 		} else {
// 			fmt.Println("🎉  YOU WON!")
// 		}
// 		return
// 	}

// 	fmt.Println("☠️  YOU LOST... Try again?")
// }
// ##########################################################	First Turn Winner	############################################################################
// ########################################################### Lucky-number-exercises/02-random-messages/solu####

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"
// )

// const (
// 	maxTurns = 5 // less is more difficult
// 	usage    = `Welcome to the Lucky Number Game! 🍀
// The program will pick %d random numbers.
// Your mission is to guess one of those numbers.
// The greater your number is, harder it gets.
// Wanna play?
// `
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	args := os.Args[1:]

// 	if len(args) != 1 {
// 		fmt.Printf(usage, maxTurns)
// 		return
// 	}

// 	guess, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		fmt.Println("Not a number.")
// 		return
// 	}

// 	if guess <= 0 {
// 		fmt.Println("Please pick a positive number.")
// 		return
// 	}

// 	for turn := 0; turn < maxTurns; turn++ {
// 		n := rand.Intn(guess) + 1

// 		if n == guess {
// 			switch rand.Intn(3) {
// 			case 0:
// 				fmt.Println("🎉  YOU WIN!")
// 			case 1:
// 				fmt.Println("🎉  YOU'RE AWESOME!")
// 			case 2:
// 				fmt.Println("🎉  PERFECT!")
// 			}
// 			return
// 		}
// 	}

// 	msg := "%s Try again?\n"

// 	switch rand.Intn(2) {
// 	case 0:
// 		fmt.Printf(msg, "☠️  YOU LOST...")
// 	case 1:
// 		fmt.Printf(msg, "☠️  JUST A BAD LUCK...")
// 	}

// }
// ########################################################### Lucky-number-exercises/02-random-messages/solu####
// ################################################################### Lucky-number-exercises/03-double-guesses/solution ######################################
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"
// )

// const (
// 	maxTurns = 5 // less is more difficult
// 	usage    = `Welcome to the Lucky Number Game! 🍀
// The program will pick %d random numbers.
// Your mission is to guess one of those numbers.
// The greater your number is, harder it gets.
// Wanna play?
// `
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	args := os.Args[1:]

// 	if len(args) < 1 {
// 		fmt.Printf(usage, maxTurns)
// 		return
// 	}

// 	guess, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		fmt.Println("Not a number.")
// 		return
// 	}

// 	var guess2 int
// 	if len(args) == 2 {
// 		guess2, err = strconv.Atoi(args[1])
// 		if err != nil {
// 			fmt.Println("Not a number.")
// 			return
// 		}
// 	}

// 	if guess <= 0 || guess2 <= 0 {
// 		fmt.Println("Please pick positive numbers.")
// 		return
// 	}

// 	min := guess
// 	if guess < guess2 {
// 		min = guess2
// 	}

// 	for turn := 0; turn < maxTurns; turn++ {
// 		n := rand.Intn(min) + 1

// 		if n == guess || n == guess2 {
// 			fmt.Println("🎉  YOU WIN!")
// 			return
// 		}
// 	}

// 	fmt.Println("☠️  YOU LOST... Try again?")
// }
// ################################################################### Lucky-number-exercises/03-double-guesses/solution ######################################

// // ###############################################  08-lucky-number-exercises/04-verbose-mode/solution/ #########################

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"
// )

// const (
// 	maxTurns = 5 // less is more difficult
// 	usage    = `Welcome to the Lucky Number Game! 🍀
// The program will pick %d random numbers.
// Your mission is to guess one of those numbers.
// The greater your number is, harder it gets.
// Wanna play?
// (Provide -v flag to see the picked numbers.)
// `
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	args := os.Args[1:]

// 	if len(args) < 1 {
// 		fmt.Printf(usage, maxTurns)
// 		return
// 	}

// 	var verbose bool

// 	if args[0] == "-v" {
// 		verbose = true
// 	}

// 	guess, err := strconv.Atoi(args[len(args)-1])
// 	if err != nil {
// 		fmt.Println("Not a number.")
// 		return
// 	}

// 	if guess <= 0 {
// 		fmt.Println("Please pick a positive number.")
// 		return
// 	}

// 	for turn := 1; turn <= maxTurns; turn++ {
// 		n := rand.Intn(guess) + 1

// 		if verbose {
// 			fmt.Printf("%d ", n)
// 		}

// 		if n == guess {
// 			if turn == 1 {
// 				fmt.Println("🥇 FIRST TIME WINNER!!!")
// 			} else {
// 				fmt.Println("🎉  YOU WON!")
// 			}
// 			return
// 		}
// 	}

// 	fmt.Println("☠️  YOU LOST... Try again?")
// }
// // ###############################################  08-lucky-number-exercises/04-verbose-mode/solution/ #########################
// ###################################################### 8-lucky-number-exercises/05-enough-picks #####################################

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"
// )

// const (
// 	maxTurns = 5 // less is more difficult
// 	usage    = `Welcome to the Lucky Number Game! 🍀
// The program will pick %d random numbers.
// Your mission is to guess one of those numbers.
// The greater your number is, harder it gets.
// Wanna play?
// `
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	args := os.Args[1:]

// 	if len(args) < 1 {
// 		fmt.Printf(usage, maxTurns)
// 		return
// 	}

// 	guess, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		fmt.Println("Not a number.")
// 		return
// 	}

// 	if guess <= 0 {
// 		fmt.Println("Please pick a positive number.")
// 		return
// 	}

// 	min := 10
// 	if guess > min {
// 		min = guess
// 	}

// 	for turn := 0; turn < maxTurns; turn++ {
// 		n := rand.Intn(min) + 1

// 		if n == guess {
// 			fmt.Println("🎉  YOU WIN!")
// 			return
// 		}
// 	}

// 	fmt.Println("☠️  YOU LOST... Try again?")
// }

// ###################################################### 8-lucky-number-exercises/05-enough-picks #####################################
// ############################## Dynamic-difficulty ##############################
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"
// )

// const (
// 	maxTurns = 5 // less is more difficult
// 	usage    = `Welcome to the Lucky Number Game! 🍀
// The program will pick %d random numbers.
// Your mission is to guess one of those numbers.
// The greater your number is, harder it gets.
// Wanna play?
// `
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	args := os.Args[1:]

// 	if len(args) != 1 {
// 		fmt.Printf(usage, maxTurns)
// 		return
// 	}

// 	guess, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		fmt.Println("Not a number.")
// 		return
// 	}

// 	if guess <= 0 {
// 		fmt.Println("Please pick a positive number.")
// 		return
// 	}

// 	var balancer int
// 	if guess > 10 {
// 		balancer = guess / 4
// 	}

// 	for turn := maxTurns + balancer; turn > 0; turn-- {
// 		n := rand.Intn(guess) + 1
// 		fmt.Printf(" > %d <", n)
// 		if n == guess {
// 			fmt.Println("🎉  YOU WIN!")
// 			return
// 		}
// 	}

// 	fmt.Println("☠️  YOU LOST... Try again?")
// }
// ############################## Dynamic-difficulty ##############################
