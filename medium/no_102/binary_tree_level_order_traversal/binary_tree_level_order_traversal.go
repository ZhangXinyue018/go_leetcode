package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	return RecordOrder(root, 0, [][]int{})
}

func RecordOrder(root *TreeNode, currentLevel int, result [][]int) [][]int{
	if root == nil {
		return result
	}
	if len(result) == currentLevel {
		result = append(result, []int{})
	}
	result[currentLevel] = append(result[currentLevel], root.Val)
	result = RecordOrder(root.Left, currentLevel+1, result)
	result = RecordOrder(root.Right, currentLevel+1, result)
	return result
}
