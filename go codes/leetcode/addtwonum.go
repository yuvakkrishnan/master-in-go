package main

// ListNode is node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{Val: -1}
	cur := &head
	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		val := sum % 10
		carry = sum / 10
		node := ListNode{Val: val}
		cur.Next = &node
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if carry == 0 {
		if l1 != nil {
			cur.Next = l1
		}
		if l2 != nil {
			cur.Next = l2
		}
		return head.Next
	}
	for l1 != nil {
		sum := l1.Val + carry
		val := sum % 10
		carry = sum / 10
		node := ListNode{Val: val}
		cur.Next = &node
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + carry
		val := sum % 10
		carry = sum / 10
		node := ListNode{Val: val}
		cur.Next = &node
		cur = cur.Next
		l2 = l2.Next
	}
	for carry != 0 {
		val := carry % 10
		carry = carry / 10
		node := ListNode{Val: val}
		cur.Next = &node
		cur = cur.Next
	}
	return head.Next
}

func main() {

}
