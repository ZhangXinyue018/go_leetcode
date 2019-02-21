package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	testNode := &ListNode{
		Val:  0,
		Next: nil,
	}

	for head != nil {
		if head.Next != nil && head.Next == testNode {
			return true
		}
		temp := head
		head = head.Next
		temp.Next = testNode
	}
	return false
}

func hasCycleBetter(head *ListNode) bool {
	slow, fast := head, head
	for {
		slow = next(slow)
		fast = next(next(fast))
		if slow == nil || fast == nil{
			return false
		}
		if slow == fast{
			return true
		}
	}
}

func next(n *ListNode) *ListNode {
	if n == nil {
		return nil
	}
	return n.Next
}
