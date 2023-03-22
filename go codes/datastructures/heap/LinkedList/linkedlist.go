package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (list *LinkedList) Add(value int) {
	node := &Node{value: value}

	if list.head == nil {
		list.head = node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}

	list.size++
}

func (list *LinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := &LinkedList{}
	list.Add(1)
	list.Add(5)
	list.Add(3)
	list.Print()
}
