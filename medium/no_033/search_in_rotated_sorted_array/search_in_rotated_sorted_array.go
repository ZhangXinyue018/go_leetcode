package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if target == nums[0] {
		return 0
	} else if target > nums[0] {
		for i := 1; i < len(nums) && nums[i] <= target; i++ {
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := len(nums) - 1; i > 0 && nums[i] >= target; i-- {
			if nums[i] == target {
				return i
			}
		}
	}
	return -1
}

func searchBetter(nums []int, target int) int {
	return DetailSearch(nums, 0, len(nums)-1, target)
}

func DetailSearch(nums []int, fromIndex int, toIndex int, target int) int {
	if fromIndex == toIndex && nums[fromIndex] == target {
		return fromIndex
	} else if fromIndex >= toIndex {
		return -1
	}

	mid := (fromIndex + toIndex + 1) / 2
	if target == nums[fromIndex] {
		return fromIndex
	} else if target == nums[mid] {
		return mid
	}
	if nums[mid] > nums[fromIndex] {
		if target < nums[fromIndex] || target > nums[mid] {
			return DetailSearch(nums, mid+1, toIndex, target)
		} else {
			return DetailSearch(nums, fromIndex+1, mid-1, target)
		}
	} else {
		if target < nums[fromIndex] && target > nums[mid] {
			return DetailSearch(nums, mid+1, toIndex, target)
		} else {
			return DetailSearch(nums, fromIndex+1, mid-1, target)
		}
	}
}
