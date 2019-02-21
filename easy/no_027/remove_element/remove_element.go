package remove_element

func RemoveElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		for left < len(nums) && nums[left] != val {
			left++
		}
		for right >= 0 && nums[right] == val {
			right--
		}

		if left < right {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
			right--
		}
	}
	return left
}
