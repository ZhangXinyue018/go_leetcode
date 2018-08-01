package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// todo: to be implemented
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cin := 0
	var headNode = &ListNode{Val: 0, Next: nil,}
	currNode := headNode
	for l1 != nil || l2 != nil {
		var result int
		if l1 == nil {
			result = l2.Val + cin
			l2 = l2.Next
		} else if l2 == nil {
			result = l1.Val + cin
			l1 = l1.Next
		} else {
			result = l1.Val + l2.Val + cin
			l1 = l1.Next
			l2 = l2.Next
		}
		newNode := &ListNode{Val: result % 10, Next: nil}
		currNode.Next = newNode
		currNode = newNode
		cin = result / 10
	}
	if cin != 0 {
		currNode.Next = &ListNode{Val: cin, Next: nil}
	}
	return headNode.Next
}
