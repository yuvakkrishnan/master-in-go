//This program done by myself

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func twoSort(a, b []int) []int {
	lenArr := make([]int, len(a)+len(b))
	i := 0
	j := 0
	k := 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			lenArr[k] = a[i]
			i++

		} else {
			lenArr[k] = b[j]
			j++
		}
		k++

	}

	for i < len(a) {
		lenArr[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		lenArr[k] = b[j]
		j++
		k++
	}
	return lenArr
}
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// fmt.Println(mid)
	// fmt.Println(len(left))
	// fmt.Println(len(right))

	return twoSort(left, right)
}

func removeDuplicates(dup []int) []int {
	dupMap := make(map[int]bool)
	duplicate := []int{}
	for _, v := range dup {
		if !dupMap[v] {
			dupMap[v] = true
			duplicate = append(duplicate, v)
		}
	}
	result := duplicate[:10]
	return mergeSort(result)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mapHash := make(map[int]int)
	for i := 0; i < 20; i++ {
		var num int
		for num == 0 {
			num = rand.Intn(20) + 1
			mapHash[i] = num
		}
		//fmt.Println(mapHash)
	}
	slcArr := []int{}
	for _, value := range mapHash {
		slcArr = append(slcArr, value)
	}

	r := removeDuplicates(slcArr)
	fmt.Println(len(r))
	fmt.Println(r)

}
