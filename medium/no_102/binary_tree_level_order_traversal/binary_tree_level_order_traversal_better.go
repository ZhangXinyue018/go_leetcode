package binary_tree_level_order_traversal

func levelOrderBetter(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	mark1 := []*TreeNode{root}
	mark2 := []*TreeNode{}
	result := [][]int{{root.Val}}
	for len(mark1) > 0 {
		for _, tempNode := range mark1 {
			if (*tempNode).Left != nil {
				mark2 = append(mark2, (*tempNode).Left)
			}
			if (*tempNode).Right != nil {
				mark2 = append(mark2, (*tempNode).Right)
			}
		}
		if len(mark2) <= 0 {
			break
		}
		layerResult := []int{}
		for _, tempNode := range mark2 {
			layerResult = append(layerResult, (*tempNode).Val)
		}
		result = append(result, layerResult)
		mark1 = mark2
		mark2 = []*TreeNode{}
	}
	return result
}
