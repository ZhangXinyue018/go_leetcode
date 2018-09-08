package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	headNode := &ListNode{Val: 0, Next: nil}
	currNode := headNode
	for l1 != nil || l2 != nil {
		var tempVal int
		if l1 == nil || (l2 != nil && l2.Val < l1.Val) {
			tempVal = l2.Val
			l2 = l2.Next
		} else {
			tempVal = l1.Val
			l1 = l1.Next
		}
		currNode.Next = &ListNode{Val: tempVal, Next: nil}
		currNode = currNode.Next
	}
	return headNode.Next
}

func main() {
	printList(mergeTwoLists(
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}))
}

func printList(node *ListNode) () {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
