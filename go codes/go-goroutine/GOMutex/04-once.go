package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var config map[string]string

func loadConfig() {
	fmt.Println("Loading config....")
	config = map[string]string{
		"Rt3bic": "value1", "92NIRJ": "value2", "074JSQ": "value3", "X6CUFI": "value4", "Ilx58T": "value5", "I8ZDtW": "value5", "pdIWMk": "value6", "0vXLEA": "value7", "oDdZ4X": "value8",
	}
}

func getConfig(key string) string {
	once.Do(loadConfig)
	return config[key]
}

func main() {
	fmt.Println(getConfig("074JSQ"))
}
