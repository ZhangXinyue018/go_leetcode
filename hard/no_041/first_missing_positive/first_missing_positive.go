package first_missing_positive

func firstMissingPositive(nums []int) int {
	temp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= len(nums) && nums[i] > 0 {
			temp[nums[i]] = 1
		}
	}
	for i := 1; i < len(temp); i++ {
		if temp[i] == 0 {
			return i
		}
	}
	return len(temp)
}
