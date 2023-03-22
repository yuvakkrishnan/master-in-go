package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	usage = `[Your Name]`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(usage)
		return
	}
	name := args[0]
	moods := [...]string{
		"Feelsad",
		"Terrible",
		"Happy",
		"Cheerful",
		"Reflective",
		"Gloomy",
		"Humorous",
		"Melancholy",
		"Idyllic",
		"Whimsical",
		"Romantic",
	}
	n := rand.Intn(len(moods))
	fmt.Printf("%s feels %s\n", name, moods[n])

}
