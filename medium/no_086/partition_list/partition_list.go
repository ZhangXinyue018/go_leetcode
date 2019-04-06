package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	remark := head
	temp := remark.Next
	pre := head
	for temp != nil {
		if temp.Val < x {
			if temp != remark.Next {
				remarkNext := remark.Next
				remark.Next = temp
				pre.Next = temp.Next
				temp.Next = remarkNext
				temp = pre
			}
			remark = remark.Next
		}
		pre = temp
		temp = temp.Next
	}
	return head.Next
}
