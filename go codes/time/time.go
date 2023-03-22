package main

import (
	"fmt"
	"time"
)

func main() {
	var timeNow time.Time = time.Now()
	var yearNow int = timeNow.Year()
	fmt.Println(yearNow)
}
