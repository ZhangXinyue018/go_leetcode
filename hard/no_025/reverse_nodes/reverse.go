package reverse_nodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return reverseK(head, k)
}

func reverseK(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	pointer := head
	first := head
	count := 0
	for count < k-1 && pointer != nil {
		pointer = pointer.Next
		count++
	}
	if count < k-1 || pointer == nil {
		return head
	}
	next := reverseK(pointer.Next, k)
	pointer.Next = nil
	for first != nil {
		temp := first.Next
		first.Next = next
		next = first
		first = temp
	}
	return next
}
