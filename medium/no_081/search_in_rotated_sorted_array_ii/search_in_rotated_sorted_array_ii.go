package search_in_rotated_sorted_array_ii

func search(nums []int, target int) bool {
	return DetailSearch(nums, 0, len(nums)-1, target)
}

func DetailSearch(nums []int, left, right, target int) bool {
	if left > right {
		return false
	}
	mid := (left + right) / 2
	if target == nums[left] || target == nums[mid] || target == nums[right] {
		return true
	}
	i := left + 1
	for ; i <= right && nums[i] == nums[left]; i++ {
	}
	return DetailSearch(nums, i, right, target)

	if nums[mid] < nums[right] && target > nums[mid] && target < nums[right] {
		return DetailSearch(nums, mid+1, right-1, target)
	} else if nums[left] < nums[mid] && target > nums[left] && target < nums[mid] {
		return DetailSearch(nums, left+1, mid-1, target)
	} else if nums[mid] > nums[right] {
		return DetailSearch(nums, mid+1, right-1, target)
	} else {
		return DetailSearch(nums, left+1, mid-1, target)
	}
}
