package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	// fmt.Println("Original:", items)

	mid := len(items) / 2
	fmt.Println(mid)

	smid := items[mid-1 : mid+2]
	fmt.Println("Test Area     :", smid)

	// sorting the smid will affect the items
	// as well. their backing array is the same.
	sort.Strings(smid)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
