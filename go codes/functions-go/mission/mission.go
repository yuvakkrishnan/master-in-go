package main

import "fmt"

func main() {
	printRoomDetails(220, 2, 2)
}

// display information about a room
func printRoomDetails(roomNumber, size, nights int) {
	nightText := "nights"
	if nights == 1 {
		nightText = "nights"
	}

	fmt.Println(roomNumber, ":", size, "people /", nights, nightText)

}

// retrieve occupancyLevel from an occupancyRate
// From 0% to 30% occupancy rate return Low
// From 30% to 60% occupancy rate return Medium
// From 60% to 100% occupancy rate return High

func occupancyLevel(occupancyRate float32) string {

	switch {
	case occupancyRate > 70:
		return "High"
	case occupancyRate > 20:
		return "Medium"
	default:
		return "Low"
	}

}

// compute the hotel occupancy rate
// return a percentage
// ex : 14,43 => 14,43%

func occupancyRate(roomsOccupied int, totalRooms int) float32 {
	return (float32(roomsOccupied) / float32(totalRooms))
}
