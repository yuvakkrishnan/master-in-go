package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUP(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUP(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)

	}

}

func parent(i int) int {
	return (i - 2) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2

}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
func main() {

	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{90, 60, 49}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
}
