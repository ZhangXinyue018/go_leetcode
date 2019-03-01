package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	_, result := InOrderTranversal(root)
	return result
}

func InOrderTranversal(root *TreeNode) ([]int, bool) {
	if root == nil {
		return []int{}, true
	}
	resultLeft, isInOrder := InOrderTranversal(root.Left)
	if !isInOrder {
		return []int{}, false
	}
	resultLeft, isInOrder = CheckAndAppend(resultLeft, []int{root.Val})
	if !isInOrder {
		return []int{}, false
	}
	resultRight, isInOrder := InOrderTranversal(root.Right)
	if !isInOrder {
		return []int{}, false
	}

	finalResult, isInOrder := CheckAndAppend(resultLeft, resultRight)
	if !isInOrder {
		return []int{}, false
	}
	return finalResult, true

}

func CheckAndAppend(result []int, nums []int) ([]int, bool) {
	if len(result) <= 0 || len(nums) <= 0 {
		return append(result, nums...), true
	} else if result[len(result)-1] >= nums[0] {
		return []int{}, false
	} else {
		return append(result, nums...), true
	}
}
