package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	var head *Node            // Declare head of linked list as nil
	head = &Node{value: 1}    // Create first node in the list
	second := &Node{value: 2} // Create second node in the list
	head.next = second        // Link the first node to the second node

	AddNode(head, 3) // Add a third node to the end of the list

	PrintList(head) // Print the contents of the linked list
}

func AddNode(head *Node, value int) {
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{value: value} // Create a new node and link it to the end of the list
}

func PrintList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}
