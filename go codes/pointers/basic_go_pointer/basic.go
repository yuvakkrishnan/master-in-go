package main

import "fmt"

func array() {
	nums := [...]int{1, 2, 3}
	fmt.Println("Arrays OG nums  :", nums)

	incrByPtr(&nums)
	fmt.Println("Arrays incrByPtr  :", nums)
	//fmt.Printf("Arrays :%p\n", &nums)

}

func incr(nums [3]int) {
	fmt.Println(nums)

	fmt.Printf("Arrays :%p\n ", &nums)

	for i := range nums {
		nums[i]++
	}

}
func incrByPtr(nums *[3]int) {
	fmt.Println(nums)

	fmt.Printf("Arrays :%p\n ", &nums)

	for i := range nums {
		nums[i]++
	}

}

func main() {
	array()

}
