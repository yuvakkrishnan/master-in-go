package main

import (
	"fmt"
	"sync"
)

func main() {
	var b sync.Mutex

	b.Lock()
	b.Lock()
	fmt.Println("This never executes as we are in deadlock")
}
