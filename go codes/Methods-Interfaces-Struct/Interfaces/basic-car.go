package main

import "fmt"

type vehicle interface {
	accelrate()
}

type car struct {
	model, color string
}

func foo(v vehicle) {
	fmt.Println(v)
}
func (c car) accelrate() {
	fmt.Println("Accelrating??")
}

type audi struct {
	model, color string
	speed        int
}

func (a audi) accelrate() {
	fmt.Println("I am accelrating fast !!")
}

func main() {
	c1 := car{model: "Benze", color: "red"}
	d1 := audi{model: "Audi R1", color: "Blue", speed: 210}
	c1.accelrate()
	d1.accelrate()
	foo(c1)
	foo(d1)

}
