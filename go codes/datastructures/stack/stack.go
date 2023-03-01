package main

import "fmt"

type Stack []int

func (stack *Stack) push(val int) {
	*stack = append(*stack, val)
}

func (stack *Stack) pop() int {
	if len(*stack) == 0 {
		return -1
	}
	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return top

}
func main() {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)

	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
