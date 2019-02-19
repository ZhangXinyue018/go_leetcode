package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	resultPointer := head
	floatingPointer := head.Next

	tempValue := resultPointer.Val

	for floatingPointer != nil {
		if floatingPointer.Val != tempValue {
			resultPointer.Next.Val = floatingPointer.Val
			resultPointer = resultPointer.Next
			tempValue = resultPointer.Val
		}
		floatingPointer = floatingPointer.Next
	}

	resultPointer.Next = nil
	return head
}
