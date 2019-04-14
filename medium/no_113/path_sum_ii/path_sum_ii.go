package path_sum_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	finalResult := [][]int{}
	FinalPathSum(root, sum, []int{}, &finalResult)
	return finalResult
}

func FinalPathSum(root *TreeNode, sum int, pre []int, result *[][]int) {
	if root == nil {
		return
	}
	newPre := make([]int, len(pre))
	copy(newPre, pre)
	sum = sum - root.Val
	newPre = append(newPre, root.Val)
	if root.Left == nil && root.Right == nil && sum == 0 {
		*result = append(*result, newPre)
	}
	if root.Left != nil {
		FinalPathSum(root.Left, sum, newPre, result)
	}
	if root.Right != nil {
		FinalPathSum(root.Right, sum, newPre, result)
	}
}
