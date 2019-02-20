package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	result := CheckTreeAndReturnDepth(root)
	return result != -1
}

func CheckTreeAndReturnDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := CheckTreeAndReturnDepth(root.Left)
	if leftDepth == -1 {
		return -1
	}

	rightDepth := CheckTreeAndReturnDepth(root.Right)
	if rightDepth == -1 {
		return -1
	}

	if leftDepth >= rightDepth {
		if leftDepth-rightDepth > 1 {
			return -1
		} else {
			return leftDepth + 1
		}
	} else {
		if rightDepth-leftDepth > 1 {
			return -1
		} else {
			return rightDepth + 1
		}
	}
}
