package main

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
	size int
}

func AddLinkedList(head *LinkedList, v int) {
	current := &Node{value: v}
	for current.next != nil {
		current = current.next
	}
	current.next = &Node{value: v}

}

func DeletingLinkedList(head *Node, value int) *Node {
	current := head
	if current.value == value {
		return current.next
	}
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return head
		}
		current = current.next
	}
	return head

}

func main() {

}
