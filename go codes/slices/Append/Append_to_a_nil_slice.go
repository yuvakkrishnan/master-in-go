package main

import (
	"fmt"
	"time"
)

// My code
// func main() {
// 	var (
// 		Pizza      []string
// 		Departure  []time.Time
// 		Graduation []int
// 		lights     []bool
// 	)

// 	Pizza = append(Pizza, "x", "y", "z")
// 	now := time.Now()
// 	Departure = append(Departure,
// 		now,
// 		now.Add(time.Hour*24))
// 	Graduation = append(Graduation, 2017, 2018, 2019, 2020)
// 	lights = append(lights, true, false, true, true, true)

// 	fmt.Println(Pizza)
// 	fmt.Println(Departure)

// 	fmt.Println(Graduation)

// 	fmt.Println(lights)

// }
func main() {
	// ------------------------------------------------------------
	// Create nil slices
	// ------------------------------------------------------------

	// Pizza toppings
	var pizza []string

	// Departure times
	var departures []time.Time

	// Student graduation years
	var grads []int

	// On/off states of lights in a room
	var lights []bool

	// ------------------------------------------------------------
	// Append them some elements (use your creativity!)
	// ------------------------------------------------------------
	pizza = append(pizza, "pepperoni", "onions", "extra cheese")

	now := time.Now()
	departures = append(departures,
		now,
		now.Add(time.Hour*24), // 24 hours after `now`
		now.Add(time.Hour*48)) // 48 hours after `now`

	grads = append(grads, 1998, 2005, 2018)

	lights = append(lights, true, false, true)

	// ------------------------------------------------------------
	// Print the slices
	// ------------------------------------------------------------

	fmt.Printf("pizza       : %s\n", pizza)
	fmt.Printf("\ngraduations : %d\n", grads)
	fmt.Printf("\ndepartures  : %s\n", departures)
	fmt.Printf("\nlights      : %t\n", lights)
}
