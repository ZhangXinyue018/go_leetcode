package selection

func Selection(nums []int, targetInt int) (int) {
	return DetailedSelection(nums, 0, len(nums)-1, targetInt)
}

func DetailedSelection(nums []int, leftIndex int, rightIndex int, targetIndex int) int {
	if leftIndex == rightIndex {
		return nums[rightIndex]
	}

	temp := nums[leftIndex]
	leftPointer := leftIndex + 1
	rightPointer := rightIndex

	for leftPointer <= rightPointer {
		for leftPointer < len(nums) && nums[leftPointer] < temp {
			leftPointer++
		}
		for rightPointer > leftIndex && nums[rightPointer] > temp {
			rightPointer--
		}
		if leftPointer < rightPointer {
			nums[leftPointer], nums[rightPointer] = nums[rightPointer], nums[leftPointer]
		}
	}
	leftCount := rightPointer - leftIndex
	if targetIndex == leftCount+1 {
		return nums[leftIndex]
	}
	nums[leftIndex], nums[rightPointer] = nums[rightPointer], nums[leftIndex]
	if targetIndex < leftCount+1 {
		return DetailedSelection(nums, leftIndex, rightPointer-1, targetIndex)
	} else {
		return DetailedSelection(nums, rightPointer+1, rightIndex, targetIndex-leftCount-1)
	}
}
