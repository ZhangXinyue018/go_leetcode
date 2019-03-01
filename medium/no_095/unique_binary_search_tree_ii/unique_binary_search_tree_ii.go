package unique_binary_search_tree_ii

import "strconv"

var (
	composeKey = make(map[string][]*TreeNode)
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return DetailedGenerateTree(1, n)
}

func DetailedGenerateTree(fromIndex int, toIndex int) []*TreeNode {
	key := strconv.Itoa(fromIndex) + "-" + strconv.Itoa(toIndex)
	if value, ok := composeKey[key]; ok {
		return value
	}

	result := []*TreeNode{}
	if fromIndex == toIndex {
		result = []*TreeNode{{Val: fromIndex}}
	} else if fromIndex > toIndex {
		result = []*TreeNode{nil}
	} else {
		for i := fromIndex; i <= toIndex; i++ {
			leftPossibilities := DetailedGenerateTree(fromIndex, i-1)
			rightPossibilities := DetailedGenerateTree(i+1, toIndex)
			for _, left := range leftPossibilities {
				for _, right := range rightPossibilities {
					root := &TreeNode{
						Val:   i,
						Left:  left,
						Right: right,
					}
					result = append(result, root)
				}
			}
		}
	}
	composeKey[key] = result
	return result
}
