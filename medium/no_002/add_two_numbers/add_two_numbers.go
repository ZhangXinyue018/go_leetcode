package add_two_numbers

// ************* Problem #2 *************

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) Result() (string) {
	resultStr := ""
	for node != nil {
		resultStr = fmt.Sprintf("%s%d", resultStr, node.Val)
		node = node.Next
	}
	return resultStr
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultHead := &ListNode{}
	temp := resultHead
	cin := 0
	for l1 != nil || l2 != nil || cin != 0 {
		if l1 != nil {
			cin = cin + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			cin = cin + l2.Val
			l2 = l2.Next
		}
		temp.Next = &ListNode{
			Val:  cin % 10,
			Next: nil,
		}
		temp = temp.Next
		cin = cin / 10
	}
	return resultHead.Next
}
