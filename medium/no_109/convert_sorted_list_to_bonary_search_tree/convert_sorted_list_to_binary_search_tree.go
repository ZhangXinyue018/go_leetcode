package convert_sorted_list_to_bonary_search_tree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	head = &ListNode{Val: 0, Next: head}
	one, two, pre := head, head, head
	for one != nil && two != nil {
		pre = one
		one = one.Next
		if two.Next == nil {
			two = two.Next
		} else {
			two = two.Next.Next
		}
	}
	pre.Next = nil
	return &TreeNode{
		Val:   one.Val,
		Left:  sortedListToBST(head.Next),
		Right: sortedListToBST(one.Next),
	}

}
