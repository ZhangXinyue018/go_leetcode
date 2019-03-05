package quick_sort

func QuickSort(nums []int) []int {
	DetailedQuickSort(nums, 0, len(nums)-1)
	return nums
}

func DetailedQuickSort(nums []int, leftIndex int, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}
	temp := nums[leftIndex]
	leftPointer := leftIndex + 1
	rightPointer := rightIndex
	for leftPointer <= rightPointer {
		for leftPointer < len(nums) && nums[leftPointer] <= temp {
			leftPointer++
			continue
		}
		for rightPointer > leftIndex && nums[rightPointer] > temp {
			rightPointer--
			continue
		}
		if leftPointer < rightPointer {
			nums[leftPointer], nums[rightPointer] = nums[rightPointer], nums[leftPointer]
		}
	}
	nums[rightPointer], nums[leftIndex] = nums[leftIndex], nums[rightPointer]
	DetailedQuickSort(nums, leftIndex, rightPointer-1)
	DetailedQuickSort(nums, leftPointer, rightIndex)
}
