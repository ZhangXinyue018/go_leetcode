package next_permutation

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	for j := len(nums) - 1; j > (len(nums)+i)/2; j-- {
		nums[j], nums[len(nums)+i-j] = nums[len(nums)+i-j], nums[j]
	}
	if i < 0{
		return
	}
	j := i + 1
	for ; j < len(nums); j++ {
		if nums[j] > nums[i] {
			break
		}
	}
	nums[i], nums[j] = nums[j], nums[i]
}
