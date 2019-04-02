package subsets

func subsets(nums []int) [][]int {
	result := &[][]int{}
	Permutation([]int{}, nums, result)
	return *result
}

func Permutation(pre []int, nums []int, res *[][]int) {
	if len(nums) == 0 {
		(*res) = append(*res, pre)
		return
	}
	Permutation(pre, nums[1:], res)
	Permutation(append([]int{nums[0]}, pre...), nums[1:], res)
}
