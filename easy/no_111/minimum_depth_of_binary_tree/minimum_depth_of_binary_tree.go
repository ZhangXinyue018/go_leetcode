package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		leftDepth := minDepth(root.Left)
		rightDepth := minDepth(root.Right)
		if leftDepth < rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}
