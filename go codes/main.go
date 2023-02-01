package main

import (
	"fmt"
	"math"
)

func main() {
	dim1 := -2.5
	dim2 := 3.14
	dim3 := 1.0

	// Calculate the distance from zero for each dimension
	d1 := math.Abs(dim1)
	d2 := math.Abs(dim2)
	d3 := math.Abs(dim3)

	// Find the dimension closest to zero
	var nearest float64
	if d1 < d2 && d1 < d3 {
		nearest = dim1
	} else if d2 < d1 && d2 < d3 {
		nearest = dim2
	} else {
		nearest = dim3
	}

	fmt.Println("Nearest dimension to zero:", nearest)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	dimensions := []int{1, 3, 5, 4, 2, 6, 8}
	min := math.MaxInt32
	nearest := 0
	for _, dimension := range dimensions {
		if abs := int(math.Abs(float64(dimension))); abs < min {
			min = abs
			nearest = dimension
		}
	}
	fmt.Println("The nearest dimension to zero is:", nearest)
}
