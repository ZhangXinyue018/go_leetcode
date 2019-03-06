package stack

type Stack struct {
	Items []int
}

func (stack *Stack) Push(num int) {
	stack.Items = append(stack.Items, num)
}

func (stack *Stack) Pop() int {
	temp := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[:len(stack.Items)-1]
	return temp
}

func (stack *Stack) Top() int {
	return stack.Items[len(stack.Items)]
}
