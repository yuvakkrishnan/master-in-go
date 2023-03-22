package main

import (
	"fmt"
	"math"
	"sync"
)

// Point represents a point in the galaxy
type Point struct {
	x, y float64
}

// Distance calculates the distance between two points
func (p Point) Distance(q Point) float64 {
	return math.Sqrt((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y))
}

func main() {
	const K = 5           // number of nearest points to find
	const chunkSize = 100 // size of data chunks sent to each Go routine

	// data represents the points in the galaxy
	data := []Point{
		{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10},
		{11, 12}, {13, 14}, {15, 16}, {17, 18}, {19, 20},
		// ...
	}

	var wg sync.WaitGroup
	closest := make(chan Point, K) // channel to receive closest points from Go routines

	// Split the data into chunks and send each chunk to a separate Go routine
	for i := 0; i < len(data); i += chunkSize {
		wg.Add(1)
		go func(data []Point, earth Point) {
			defer wg.Done()
			var closestPoint Point
			minDistance := math.MaxFloat64
			for _, point := range data {
				distance := point.Distance(earth)
				if distance < minDistance {
					minDistance = distance
					closestPoint = point
				}
			}
			closest <- closestPoint
		}(data[i:min(i+chunkSize, len(data))], data[0]) // first earth point is data[0]
	}

	// Listen to the channel and keep track of the closest points found
	var nearest []Point
	for i := 0; i < K; i++ {
		point := <-closest
		nearest = append(nearest, point)
		data = remove(data, point) // update data to exclude the closest point
	}

	wg.Wait()
	fmt.Println("Nearest points:", nearest)
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// remove removes the specified point from the data
func remove(data []Point, point Point) []Point {
	var result []Point
	for _, p := range data {
		if p != point {
			result = append(result, p)
		}
	}
	return result
}
