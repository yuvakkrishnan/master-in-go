package main

import "fmt"

func wordrsProduce(word chan []string) []string {
	outputCH := <-word
	outputCH = append(outputCH, "We Programmers are Indians")
	result := []string{}
	for _, v := range outputCH {
		result = append(result, v)
	}
	return result
}
func main() {
	inputCh := make(chan []string)
	var messages = []string{"I am a Indian", "Our Capital is delhi.", "Our PM is Mr.Modi", "I love my country", "Our national flag is Niraga", "citys", "Goa", "Nagpur"}
	go func() {
		inputCh <- messages
	}()
	finalResult := wordrsProduce(inputCh)
	fmt.Println(finalResult)
}
