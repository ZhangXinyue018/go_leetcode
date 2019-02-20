package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	} else {
		return hasPathSum(root.Right, sum-root.Val) || hasPathSum(root.Left, sum-root.Val)
	}
}
