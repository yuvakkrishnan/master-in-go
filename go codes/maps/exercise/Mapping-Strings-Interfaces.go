package main

import "fmt"

type Service interface {
	sayHi()
}

type MyService struct{}

func (s MyService) sayHi() {
	fmt.Println("Hi")
}

type SecondService struct {
}

func (s SecondService) sayHi() {
	fmt.Println("Hello From the 2nd Service")
}

func main() {
	fmt.Println("Printing Map")
	interfaceMap := make(map[string]Service, 0)
	interfaceMap["ServiceID-1"] = MyService{}
	interfaceMap["ServiceID-2"] = SecondService{}

	//interfaceMap["ServiceID-2"].sayHi()
	for key, service := range interfaceMap {
		fmt.Println(key)
		service.sayHi()
	}
}
