// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {

// 	args := os.Args
// 	if len(args) != 2 {
// 		fmt.Printf("Give me the magnitude of the earthquake.")
// 		return
// 	}
// 	earthQuake, err := strconv.ParseFloat(args[1], 64)
// 	if err != nil {
// 		fmt.Println("i couldn't get that sorry.")
// 		return
// 	}
// 	switch {
// 	case earthQuake < 2.0:
// 		fmt.Println("micro")

// 	case earthQuake >= 2.0 && earthQuake < 3.0:
// 		fmt.Println("very minor")
// 	case earthQuake >= 3.0 && earthQuake < 4.0:
// 		fmt.Println("minor")
// 	case earthQuake >= 4.0 && earthQuake < 5.0:
// 		fmt.Println("light")

// 	case earthQuake >= 5.0 && earthQuake < 6.0:
// 		fmt.Println(" moderate")
// 	case earthQuake >= 6.0 && earthQuake < 7.0:
// 		fmt.Println("strong")
// 	case earthQuake >= 7.0 && earthQuake < 8.0:
// 		fmt.Println("major")

// 	case earthQuake >= 8.0 && earthQuake < 10:
// 		fmt.Println("great")
// 	default:
// 		fmt.Println("massive")
// 	}

// }

// *************************************************************************Username && Password
// package main

// import (
// 	"fmt"
// 	"os"
// )

// const (
// 	usage       = "Usage: [username] [password]"
// 	errUser     = "Access denied for %q.\n"
// 	errPwd      = "Invalid password for %q.\n"
// 	accessOK    = "Access granted to %q.\n"
// 	user, user2 = "jack", "inanc"
// 	pass, pass2 = "1888", "1879"
// )

// func main() {

// 	args := os.Args
// 	if len(args) != 3 {
// 		fmt.Printf("Usage")
// 		return
// 	}
// 	u, p := args[1], args[2]
// 	switch {
// 	case u != user && u != user2:
// 		fmt.Printf(errUser, u)

// 	case u == user && p == pass:
// 		fmt.Printf(accessOK, u)
// 	case u == user2 && p == pass2:
// 		fmt.Printf("AccessOk", u)

// 	default:
// 		fmt.Printf(errPwd, u)
// 	}

// }
package main
