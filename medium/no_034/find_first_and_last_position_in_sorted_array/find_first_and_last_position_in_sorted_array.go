package find_first_and_last_position_in_sorted_array

func searchRange(nums []int, target int) []int {
	return DetailSearchRange(nums, 0, len(nums)-1, target)
}

func DetailSearchRange(nums []int, fromIndex int, toIndex int, target int) []int {
	if fromIndex > toIndex {
		return []int{-1, -1}
	} else if nums[fromIndex] == nums[toIndex] && nums[fromIndex] == target {
		return []int{fromIndex, toIndex}
	}
	mid := (fromIndex + toIndex) / 2
	if target > nums[mid] {
		return DetailSearchRange(nums, mid+1, toIndex, target)
	} else if target < nums[mid] {
		return DetailSearchRange(nums, fromIndex, mid-1, target)
	} else {
		left, right := mid, mid
		for ; left >= fromIndex; left-- {
			if nums[left] != target {
				break
			}
		}
		left++
		for ; right <= toIndex; right++ {
			if nums[right] != target {
				break
			}
		}
		right--
		return []int{left, right}
	}
}
