package binary_tree_zigzag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
		if len(mark2) == 0 {
			break
		}
		tempResult := []int{}
		if len(result)%2 == 0 {
			for i := 0; i < len(mark2); i++ {
				tempResult = append(tempResult, mark2[i].Val)
			}
		} else {
			for i := len(mark2) - 1; i >= 0; i-- {
				tempResult = append(tempResult, mark2[i].Val)
			}
		}
		result = append(result, tempResult)
		mark1 = mark2
		mark2 = []*TreeNode{}
	}
	return result
}
