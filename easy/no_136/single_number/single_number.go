package single_number

func singleNumber(nums []int) int {
	return detailFind(nums, 0, len(nums)-1)
}

func detailFind(nums []int, fromIndex int, toIndex int) int {
	if fromIndex == toIndex {
		return nums[fromIndex]
	}
	for i := fromIndex + 1; i <= toIndex; i++ {
		if nums[i] == nums[fromIndex] {
			temp := nums[i]
			nums[i] = nums[toIndex]
			nums[toIndex] = temp
			return detailFind(nums, fromIndex+1, toIndex-1)
		}
	}
	return nums[fromIndex]
}

func singleNumberBetter(nums []int) int {
	result := 0
	for _, num := range nums{
		result = result ^ num
	}
	return result
}
