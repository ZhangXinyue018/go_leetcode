package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	fp := head
	sp := head
	i := 1
	for i <= k {
		if fp.Next == nil {
			k = k - (k/i)*i
			fp = head
			i = 1
		} else {
			fp = fp.Next
			i++
		}
	}
	if fp == head {
		return head
	}
	for fp.Next != nil {
		fp = fp.Next
		sp = sp.Next
	}
	fp.Next = head
	temp := sp.Next
	sp.Next = nil
	return temp
}
