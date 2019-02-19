package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return DetailedIsSymmetric(root.Left, root.Right)
}

func DetailedIsSymmetric(leftRoot *TreeNode, rightRoot *TreeNode) bool {
	if leftRoot == nil && rightRoot == nil {
		return true
	} else if leftRoot == nil || rightRoot == nil || leftRoot.Val != rightRoot.Val {
		return false
	}

	return DetailedIsSymmetric(leftRoot.Left, rightRoot.Right) && DetailedIsSymmetric(leftRoot.Right, rightRoot.Left)
}
