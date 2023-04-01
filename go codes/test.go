package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)

	for _ = range ticker.C {
		fmt.Println("Tick-TOck")

	}
}
