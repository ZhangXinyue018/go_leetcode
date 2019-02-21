package intersection_of_two_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pointerA := headA
	pointerB := headB
	for pointerA != pointerB {
		if pointerA == nil {
			pointerA = headB
		} else {
			pointerA = pointerA.Next
		}
		if pointerB == nil {
			pointerB = headA
		} else {
			pointerB = pointerB.Next
		}
	}
	return pointerA
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil {
		return nil
	}
	temp := headA
	for temp != nil {
		pointer := headB
		for pointer != nil {
			if pointer == temp {
				return temp
			}
			pointer = pointer.Next
		}
		temp = temp.Next
	}
	return nil
}
