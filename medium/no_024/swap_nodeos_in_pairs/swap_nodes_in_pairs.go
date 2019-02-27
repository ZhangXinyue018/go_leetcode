package swap_nodeos_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	head = &ListNode{0, head}
	temp := head
	for temp != nil && temp.Next != nil && temp.Next.Next != nil {
		firstNode := temp.Next
		secondNode := temp.Next.Next

		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		temp.Next = secondNode

		temp = temp.Next.Next
	}
	return head.Next
}
