package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return FormatTree(nums, 0, len(nums)-1)
}

func FormatTree(nums []int, fromIndex int, toIndex int) *TreeNode {
	if fromIndex > toIndex {
		return nil
	}

	middleIndex := (fromIndex + toIndex) / 2
	return &TreeNode{
		Val:   nums[middleIndex],
		Left:  FormatTree(nums, fromIndex, middleIndex-1),
		Right: FormatTree(nums, middleIndex+1, toIndex),
	}
}