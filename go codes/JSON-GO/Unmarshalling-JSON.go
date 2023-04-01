package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type SensorReading struct {
	Meter_id         string    `json:"meter_id"`
	Reading          float64   `json:"reading"`
	Unit             string    `json:"unit"`
	Timestamp        time.Time `json:"timestamp"`
	Ok               string    `json:"status"`
	Location         Location  `json:"location"`
	Manufacturer     string    `json:"manufacturer"`
	Serial_number    string    `json:"serial_number"`
	Software_version string    `json:"software_version"`
	Battery_level    int64     `json:"battery_level"`
	Signal_strength  int64     `json:"signal_strength"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
}

func main() {
	// Open the JSON file
	jsFile, err := os.Open("gas_meter.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer jsFile.Close()

	// Parse the JSON file
	var reading SensorReading
	err = json.NewDecoder(jsFile).Decode(&reading)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the result
	fmt.Println("Sensor Reading:", reading)
}
