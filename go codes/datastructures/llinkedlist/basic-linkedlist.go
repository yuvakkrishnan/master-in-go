package main

import "fmt"

type NodeDaata interface {
	nodeData(value1, value2, value3 int)
}
type Node struct {
	data int
	next *Node
}
type myNodeData struct{}

func (n *Node) nodeData(value1, value2, value3 int) {
	head := &Node{data: value1}
	second := &Node{data: value2}
	third := &Node{data: value3}

	head.next = second
	second.next = third
	current := head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	v1, v2, v3 := 1, 2, 3

	NodeDaata.nodeData(v1, v2, v3)

}
