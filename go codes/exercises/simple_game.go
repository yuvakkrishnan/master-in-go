package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	second := time.Now().Unix()
	rand.Seed(second)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 to 100.")
	fmt.Println("can you guess it ?")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make a guess: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func findNumberGame(num float64, target int) string {
// 	var guess string
// 	if num >= float64(target) {
// 		guess = "Oops. Your guess was HIGh"
// 	} else if num <= float64(target) {
// 		guess = "Oops. Your guess was LOW"
// 	} else if num == float64(target) {
// 		guess = "GOOD JOB! YOU GUESED IT !!!"
// 	}
// 	return guess

// }
// func gettingInput() float64 {
// 	fmt.Println("Guess Number : ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)

// 	}
// 	input = strings.TrimSpace(input)
// 	finInput, err := strconv.ParseFloat(input, 64)
// 	if err != nil {
// 		log.Fatal(err)

// 	}
// 	return finInput
// }
// func main() {
// 	seconds := time.Now().Unix()
// 	rand.Seed(seconds)
// 	target := rand.Intn(100) + 1
// 	count := 10
// 	for i := 0; i < count; i++ {
// 		fmt.Printf("You have %d guess left\n", 10-i)

// 		result := gettingInput()

// 		guess := findNumberGame(result, target)

// 		if guess == "GOOD JOB! YOU GUESED IT !!!" {
// 			break
// 		}
// 	}
// 	fmt.Printf("The target number was %d\n", target)
// }
