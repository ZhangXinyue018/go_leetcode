package validate_binary_search_tree

func isValidBSTBetter(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBSTBetter(root.Left) {
		return false
	}
	if !isValidBSTBetter(root.Right) {
		return false
	}
	if root.Left != nil && root.Val <= GetMax(root.Left) {
		return false
	}
	if root.Right != nil && root.Val >= GetMin(root.Right) {
		return false
	}
	return true
}

func GetMin(root *TreeNode) int {
	if root.Left == nil {
		return root.Val
	}
	return GetMin(root.Left)
}

func GetMax(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}
	return GetMax(root.Right)
}
