package binary_tree_level_order_traverdal_ii

// todo: memory usage is high, reduce later

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := map[int][]int{}
	finalLevel := AppendOrder(root, result, 0)
	var finalResult [][]int
	for i := finalLevel; i >= 0; i-- {
		if len(result[i]) > 0 {
			finalResult = append(finalResult, result[i])
		}
	}
	return finalResult
}

func AppendOrder(root *TreeNode, result map[int][]int, levelNo int) int {
	if root == nil {
		return levelNo
	}
	if _, hasValue := result[levelNo]; hasValue {
		result[levelNo] = append(result[levelNo], root.Val)
	} else {
		result[levelNo] = []int{root.Val}
	}
	leftLevelNo := AppendOrder(root.Left, result, levelNo+1)
	rightLevelNo := AppendOrder(root.Right, result, levelNo+1)
	if leftLevelNo > rightLevelNo {
		return leftLevelNo
	} else {
		return rightLevelNo
	}
}

func levelOrderBottom2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	finalLevel := AppendOrder2(root, &result, 0)

	for i := 0; i <= (finalLevel-1)/2; i++ {
		temp := result[i]
		result[i] = result[finalLevel-1-i]
		result[finalLevel-1-i] = temp
	}
	if len(result[0]) == 0 {
		return result[1:]
	} else {
		return result
	}
}

func AppendOrder2(root *TreeNode, result *[][]int, levelNo int) int {
	if root == nil {
		return levelNo
	}
	if len(*result) < levelNo+1 {
		*result = append(*result, []int{root.Val})
	} else {
		(*result)[levelNo] = append((*result)[levelNo], root.Val)
	}
	leftLevelNo := AppendOrder2(root.Left, result, levelNo+1)
	rightLevelNo := AppendOrder2(root.Right, result, levelNo+1)
	if leftLevelNo > rightLevelNo {
		return leftLevelNo
	} else {
		return rightLevelNo
	}
}
