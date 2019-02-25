package remove_nth_node_from_end_of_list

type Stack struct {
	NodeList []*ListNode
}

func (stack *Stack) Push(node *ListNode) {
	stack.NodeList = append(stack.NodeList, node)
}

func (stack *Stack) Pop() *ListNode {
	result := stack.NodeList[len(stack.NodeList)-1]
	stack.NodeList = stack.NodeList[:len(stack.NodeList)-1]
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	stack := Stack{}
	temp := head
	for temp != nil {
		stack.Push(temp)
		temp = temp.Next
	}
	if n == len(stack.NodeList) {
		return head.Next
	}

	var toBeDeleted *ListNode
	for i := n; i >= 0; i-- {
		toBeDeleted = stack.Pop()
	}
	toBeDeleted.Next = toBeDeleted.Next.Next
	return head
}
