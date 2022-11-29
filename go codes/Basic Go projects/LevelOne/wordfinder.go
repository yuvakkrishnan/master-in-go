// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const corpus = "lazy cat jumps again and again and again"

// func main() {
// 	words := strings.Fields(corpus)
// 	query := os.Args[1:]

// 	// after the inner loop breaks
// 	// this parent loop will look for the next queried word
// 	for _, q := range query {

// 		// "break" will terminate this loop
// 		for i, w := range words {
// 			if q == w {
// 				fmt.Printf("#%-2d: %q\n", i+1, w)

// 				// find the first word then break
// 				// the nested loop
// 				break
// 			}
// 		}

// 	}
// }

// ################################################ word-finder-exercises-case-insensitive/solution ################################################
package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// const corpus = "lazy cat jumps again and again and again"

// func main() {
// 	words := strings.Fields(corpus)
// 	query := os.Args[1:]

// queries:
// 	for _, q := range query {
// 		// case insensitive search
// 		q = strings.ToLower(q)

// 	search:
// 		for i, w := range words {
// 			switch q {
// 			case "and", "or", "the":
// 				break search
// 			}

// 			if q == w {
// 				fmt.Printf("#%-2d: %q\n", i+1, w)
// 				continue queries
// 			}
// 		}
// 	}
// }

// ################################################ word-finder-exercises-case-insensitive/solution ################################################

// ##############################################word-finder-exercises-path-searcher  ##################
// package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"
// )

// func main() {
// 	// Get and split the PATH environment variable

// 	// SplitList function automatically finds the
// 	// separator for the path env variable
// 	words := filepath.SplitList(os.Getenv("PATH"))

// 	// Alternative way, but above one is better:
// 	// words := strings.Split(
// 	// 	os.Getenv("PATH"),
// 	// 	string(os.PathListSeparator))

// 	query := os.Args[1:]

// 	for _, q := range query {
// 		for i, w := range words {
// 			q, w = strings.ToLower(q), strings.ToLower(w)

// 			if strings.Contains(w, q) {
// 				fmt.Printf("#%-2d: %q\n", i+1, w)
// 			}
// 		}
// 	}
// }
// ##############################################word-finder-exercises-path-searcher  ##################
