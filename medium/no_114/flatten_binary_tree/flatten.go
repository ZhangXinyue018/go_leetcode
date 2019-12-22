package flatten_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root.Right
	if root.Left != nil {
		Flatten(root.Left)
		RightMost := GetRightMost(root.Left)
		root.Right = root.Left
		RightMost.Right = temp
		root.Left = nil
	}
	Flatten(temp)
}

func GetRightMost(node *TreeNode) *TreeNode {
	temp := node
	for temp.Right != nil {
		temp = temp.Right
	}
	return temp
}

func Flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	flattenSub(root)
}

func flattenSub(root *TreeNode) *TreeNode {
	if root.Right == nil && root.Left == nil {
		return root
	}
	if root.Left == nil {
		return flattenSub(root.Right)
	} else if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return flattenSub(root.Right)
	} else {
		flattenSub(root.Left).Right = root.Right
		root.Right = root.Left
		root.Left = nil
		return flattenSub(root.Right)
	}
}
