In Go, a linked list node is typically implemented using a struct that contains the data and a pointer to the next node in the list. Here is an example of a simple node definition:

type Node struct {
    data int
    next *Node
}

In this example, the Node struct has two fields: data of type int, which holds the value of the node, and next of type *Node, which is a pointer to the next node in the list.

The next pointer allows us to link multiple nodes together to form a linked list. When we create a new node, we set its next field to nil to indicate that it is the last node in the list.

Here is an example of creating a simple linked list of three nodes:

head := &Node{data: 1}
second := &Node{data: 2}
third := &Node{data: 3}

head.next = second
second.next = third

In this example, we first create the head node with data set to 1. We then create the second and third nodes with data set to 2 and 3, respectively. Finally, we link the nodes together by setting the next pointers of the head and second nodes to the second and third nodes, respectively.

To traverse a linked list, we start at the head node and follow the next pointers until we reach the end of the list. Here is an example of printing the values of all nodes in the linked list:

current := head
for current != nil {
    fmt.Println(current.data)
    current = current.next
}
In this example, we start at the head node and print its data. We then follow the next pointer to the next node and print its data, and so on, until we reach the end of the list (when current becomes nil).

This is just a basic example of how linked lists can be implemented in Go. Depending on the use case, more complex node definitions and operations may be required.