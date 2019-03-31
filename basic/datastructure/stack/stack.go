package stack

type Stack struct {
	Items []interface{}
}

func (stack *Stack) Push(num interface{}) {
	stack.Items = append(stack.Items, num)
}

func (stack *Stack) Pop() interface{} {
	temp := stack.Items[len(stack.Items)-1]
	stack.Items = stack.Items[:len(stack.Items)-1]
	return temp
}

func (stack *Stack) Top() interface{} {
	return stack.Items[len(stack.Items)]
}
