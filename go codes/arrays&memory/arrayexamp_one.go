// package main

// import "fmt"

// // STORY:
// // Hipster's Love store publishes limited books
// // twice a year.
// //
// // The number of books they publish is fixed at 4.

// // So, let's create a 4 elements string array for the books.

// const (
// 	winter = 1
// 	summer = 3
// 	yearly = winter + summer
// )

// func main() {
// 	// var books [4]string
// 	// var books [1 + 3]string
// 	var books [yearly]string

// 	books[0] = "Kafka's Revenge"
// 	books[1] = "Stay Golden"
// 	books[2] = "Everythingship"
// 	books[3] += books[0] + " 2nd Edition"

// 	// --------------------
// 	// INDEXING
// 	// --------------------

// 	// Go compiler can catch indexing errors when constant is used
// 	// books[4] = "Neverland"

// 	// Go compiler cannot catch indexing errors when non-constant is used
// 	// i := 4
// 	// books[i] = "Neverland"

// 	// --------------------
// 	// PRINTING ARRAYS
// 	// --------------------

// 	// print the type
// 	fmt.Printf("books  : %T\n", books)

// 	// print the elements
// 	fmt.Println("books  :", books)

// 	// print the elements in quoted string
// 	fmt.Printf("books  : %q\n", books)

// 	// print the elements along with their types
// 	fmt.Printf("books  : %#v\n", books)
// }
// ###################################################################################	03-examples-2-hipsters-love-bookstore	###########################

// package main

// import "fmt"

// // STORY:
// // Hipster's Love store publishes limited books
// // twice a year.
// //
// // The number of books they publish is fixed at 4.

// // So, let's create a 4 elements string array for the books.

// const (
// 	winter = 1
// 	summer = 3
// 	yearly = winter + summer
// )

// func main() {
// 	var books [yearly]string

// 	books[0] = "Kafka's Revenge"
// 	books[1] = "Stay Golden"
// 	books[2] = "Everythingship"
// 	books[3] += books[0] + " 2nd Edition"
// 	fmt.Printf("books  : %#v\n", books)

// 	// ------------------------------------
// 	// SEASONAL BOOKS
// 	// ------------------------------------

// 	var (
// 		wBooks [winter]string
// 		sBooks [summer]string
// 	)

// 	// assign the first book as the wBook's first book
// 	wBooks[0] = books[0]

// 	// assign all the books starting from the 2nd book
// 	// to sBooks array

// 	// sBooks[0] = books[1]
// 	// sBooks[1] = books[2]
// 	// sBooks[2] = books[3]

// 	// for i := 0; i < len(sBooks); i++ {
// 	// 	sBooks[i] = books[i+1]
// 	// }

// 	for i := range sBooks {
// 		sBooks[i] = books[i+1]
// 		// changes to sBooks[i] will not be visible here.
// 		// sBooks here is a copy of the original array.
// 	}
// 	// changes to sBooks are visible here

// 	// sBooks is a copy of the original sBooks array.
// 	//
// 	// v is also a copy of the original array element.
// 	//
// 	// if you want to update the original element, use it as in the loop above.
// 	//
// 	// for _, v := range sBooks {
// 	// 	v += "won't effect"
// 	// }

// 	fmt.Printf("\nwinter : %#v\n", wBooks)
// 	fmt.Printf("\nsummer : %#v\n", sBooks)

// 	// ------------------------------------
// 	// Let's print the published books
// 	// ------------------------------------

// 	// equal to: [4]bool
// 	var published [len(books)]bool

// 	published[0] = true
// 	published[len(books)-1] = true

// 	fmt.Println("\nPublished Books:")
// 	for i, ok := range published {
// 		if ok {
// 			fmt.Printf("+ %s\n", books[i])
// 		}
// 	}
// }

// ###################################################################################	03-examples-2-hipsters-love-bookstore	###########################

// STORY:
// Hipster's Love store publishes limited books
// twice a year.
//
// The number of books they publish is fixed at 4.

// So, let's create a 4 elements string array for the books.

// func main() {
// 	// Use this only when you don't know about the elements beforehand
// 	{
// 		var books [4]string

// 		books[0] = "Kafka's Revenge"
// 		books[1] = "Stay Golden"
// 		books[2] = "Everythingship"
// 		books[3] += "Kafka's Revenge 2nd Edition"

// 		_ = books
// 	}

// 	// This is not necessary, use the short declaration syntax below
// 	{
// 		var books = [4]string{
// 			"Kafka's Revenge",
// 			"Stay Golden",
// 			"Everythingship",
// 			"Kafka's Revenge 2nd Edition",
// 		}

// 		_ = books
// 	}

// Use this if you know about the elements
// {
// 	books := [4]string{
// 		"Kafka's Revenge",
// 		"Stay Golden",
// 		"Everythingship",
// 		"Kafka's Revenge 2nd Edition",
// 	}

// 	_ = books
// }

// {
// 	// Use this if you know about the elements
// 	books := [4]string{
// 		"Kafka's Revenge",
// 		"Stay Golden",
// 	}

// 	// Uninitialized elements will be set to their zero values
// 	fmt.Printf("books  : %#v\n", books)
// }

// You can also use the ellipsis syntax
// ... equals to 4
// 	{
// 		books := [...]string{
// 			"Kafka's Revenge",
// 			"Stay Golden",
// 			"Everythingship",
// 			"Kafka's Revenge 2nd Edition",
// 		}

// 		_ = books
// // 	}
// }

// #######################################################	05-examples-3-hipsters-love-bookstore	###########################################

// package main

// import "fmt"

// // STORY:
// // Hipster's Love store publishes limited books
// // twice a year.
// //
// // The number of books they publish is fixed at 4.

// // So, let's create a 4 elements string array for the books.

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
// 	fmt.Printf("books  : %#v\n", books)
// }

// #######################################################	05-examples-3-hipsters-love-bookstore	###########################################
