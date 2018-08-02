package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// todo: to be optimized??
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cin := 0
	var headNode = &ListNode{Val: 0, Next: nil,}
	currNode := headNode
	for l1 != nil || l2 != nil || cin != 0 {
		var a, b int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		result := a + b + cin
		newNode := &ListNode{Val: result % 10, Next: nil}
		currNode.Next = newNode
		currNode = newNode
		cin = result / 10
	}
	return headNode.Next
}
