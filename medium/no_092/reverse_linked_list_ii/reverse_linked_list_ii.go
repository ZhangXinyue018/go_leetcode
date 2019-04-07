package reverse_linked_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	i := 1
	head = &ListNode{Val: 0, Next: head}
	mPre, mP, nP := head, head.Next, head.Next
	temp := head.Next
	for temp != nil && i < n {
		if i < m {
			mPre = mPre.Next
			mP = mP.Next
		}
		if i < n {
			nP = nP.Next
		}
		i++
		temp = temp.Next
	}
	pre := nP.Next
	for pre != nP {
		temp = mP.Next
		mP.Next = pre
		pre = mP
		mP = temp
	}
	mPre.Next = pre
	return head.Next
}
