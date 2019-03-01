package binary_tree_in_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return DetailInOrderTraversal(root, []int{})
}

func DetailInOrderTraversal(root *TreeNode, preResult []int) []int {
	if root == nil {
		return preResult
	}
	finalResult := DetailInOrderTraversal(root.Left, preResult)
	finalResult = append(finalResult, root.Val)
	finalResult = DetailInOrderTraversal(root.Right, finalResult)
	return finalResult
}
