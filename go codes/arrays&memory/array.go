// package main

// import "fmt"

// STORY:
// Hipster's Love store publishes limited books
// twice a year.
//
// The number of books they publish is fixed at 4.

// So, let's create a 4 elements string array for the books.

// const (
// 	winter = 1
// 	summer = 3
// 	yearly = winter + summer
// )

// func main() {
// 	books := [...]string{
// 		"Kafka's Revenge",
// 		"Stay Golden",
// 		"Everythingship",
// 		"Kafka's Revenge 2nd Edition",
// 	}

// 	var (
// 		sBooks [summer]string
// 	)

// 	for i := range sBooks {
// 		sBooks[i] = books[i+1]
// 		fmt.Println(sBooks)
// 	}

// }
// Basics of Array:
package main

import "fmt"

func main() {
	var (
		bestFriends        [3]string
		differentLocations [5]string
		dataBuffer         [5]byte
		currencyExchange   [1]float64
		upDownStatus       [4]bool
		zero               [0]byte
	)
	fmt.Printf("best friends : %#v\n", bestFriends)
	fmt.Printf("Locations : %#v\n", differentLocations)
	fmt.Printf("Data : %#v\n", dataBuffer)
	fmt.Printf("CurrencyExchange : %#v\n", currencyExchange)
	fmt.Printf("Alives : %#v\n", upDownStatus)
	fmt.Printf("Zero :%#v\n", zero)
	fmt.Println("#############################################")

	fmt.Printf("best friends : %T\n", bestFriends)
	fmt.Printf("Locations : %T\n", differentLocations)
	fmt.Printf("Data : %T\n", dataBuffer)
	fmt.Printf("CurrencyExchange : %T\n", currencyExchange)
	fmt.Printf("Alives : %T \n", upDownStatus)

}
