package main

import "fmt"

func unbufferedChanne(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println("Print of V :", v)
		fmt.Println("Print of C :", c)

		sum += v
	}
	c <- sum
}

func main() {
	ch := make(chan int)
	setOfArrary := []int{1, 2, 3, 4, 5, 6}

	go unbufferedChanne(setOfArrary[:len(setOfArrary)/2], ch)
	go unbufferedChanne(setOfArrary[len(setOfArrary)/2:], ch)
	x, y := <-ch, <-ch
	fmt.Println("Print of X :", x)
	fmt.Println("Print of Y :", y)
	fmt.Println("Print of Sum X+Y :", x+y)

}
