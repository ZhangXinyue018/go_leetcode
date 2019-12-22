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
		for leftPointer <= rightIndex && nums[leftPointer] <= temp {
			leftPointer++
		}
		for rightPointer > leftIndex && nums[rightPointer] > temp {
			rightPointer--
		}
		if leftPointer < rightPointer {
			nums[leftPointer], nums[rightPointer] = nums[rightPointer], nums[leftPointer]
		}
	}
	nums[rightPointer], nums[leftIndex] = nums[leftIndex], nums[rightPointer]
	DetailedQuickSort(nums, leftIndex, rightPointer-1)
	DetailedQuickSort(nums, leftPointer, rightIndex)
}

func QuickSort2(start, end int, nums []int) {
	if start >= end {
		return
	}
	standard := nums[start]
	left := start
	right := end
	pointer := start + 1
	for left < right {
		if nums[pointer] >= standard {
			nums[pointer], nums[right] = nums[right], nums[pointer]
			right--
		} else {
			nums[pointer], nums[left] = nums[left], nums[pointer]
			left++
			pointer++
		}
	}
	QuickSort2(start, left-1, nums)
	QuickSort2(left+1, end, nums)
}
