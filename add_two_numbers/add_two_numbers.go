package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) print() () {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cin := 0
	var headNode = &ListNode{Val: 0, Next: nil,}
	currNode := headNode
	for l1 != nil || l2 != nil || cin != 0 {
		var a, b int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		result := a + b + cin
		newNode := &ListNode{Val: result % 10, Next: nil}
		currNode.Next = newNode
		currNode = newNode
		cin = result / 10
	}
	return headNode.Next
}

func main() () {
	node11 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3,}}}
	node12 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4,}}}
	node1 := AddTwoNumbers(node11, node12)
	node1.print()
	fmt.Println()
	node21 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9,}}}
	node22 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1}}}}
	node2 := AddTwoNumbers(node21, node22)
	node2.print()
}
