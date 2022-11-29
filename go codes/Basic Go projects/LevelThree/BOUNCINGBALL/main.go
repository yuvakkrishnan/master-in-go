package main

import "fmt"

func main() {
	const (
		width     = 50
		height    = 10
		cellEmpty = ' '
		cellball  = '🏐'
	)
	var ball rune
	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
		// fmt.Println("Printing Board:  ", board)

	}
	// buf := make([]rune, 0, width*height)
	//Draw a Smiley

	// board[12][2] = true
	// board[16][2] = true

	// board[14][4] = true
	// board[10][6] = true

	// board[18][6] = true
	// board[12][7] = true

	// board[14][7] = true
	// board[16][7] = true

	for y := range board[0] {
		// fmt.Println("Printing Y:  ", y)
		for x := range board {
			ball = cellEmpty
			// fmt.Println("Printing X:  ", x)
			if board[x][y] {
				ball = cellball

			}
			fmt.Print(string(ball))
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// fmt.Print("⚽")
