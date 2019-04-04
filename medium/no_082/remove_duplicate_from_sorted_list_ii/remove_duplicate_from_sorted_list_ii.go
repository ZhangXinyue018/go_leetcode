package remove_duplicate_from_sorted_list_ii

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	tempHead := &ListNode{Val: 0, Next: head}
	preHead := &ListNode{Val: 0}
	pre := preHead
	p := tempHead.Next
	v, c := 0, 0
	for ; p != nil; p = p.Next {
		if p == head || v != p.Val {
			if c == 1 {
				pre.Next = &ListNode{Val: v}
				pre = pre.Next
			}
			v = p.Val
			c = 1
		} else {
			c++
		}
	}
	if c == 1 {
		pre.Next = &ListNode{Val: v}
	}
	return preHead.Next
}
