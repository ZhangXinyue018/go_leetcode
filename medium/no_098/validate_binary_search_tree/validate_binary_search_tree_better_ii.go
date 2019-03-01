package validate_binary_search_tree

import "math"

func isValidBSTBetter2(root *TreeNode) bool {
	return IsValidTree(root, math.MaxInt64, math.MinInt64)
}

func IsValidTree(root *TreeNode, maxValue int, minValue int) bool {
	if root == nil {
		return true
	} else if root.Val >= maxValue || root.Val <= minValue {
		return false
	}
	return IsValidTree(root.Left, root.Val, minValue) && IsValidTree(root.Right, maxValue, root.Val)
}
