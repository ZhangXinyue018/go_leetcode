package merge_sort

import "fmt"

// 二分归并法可归类为分治法

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	half := len(nums) / 2
	left := MergeSort(nums[:half])
	right := MergeSort(nums[half:])
	i, j := 0, 0
	result := []int{}
	for i < len(left) || j < len(right) {
		if i == len(left) {
			result = append(result, right[j:]...)
			break
		} else if j == len(right) {
			result = append(result, left[i:]...)
			break
		} else if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	return result
}

func MergeSortLowMemVersion(nums []int, leftIndex int, rightIndex int) {
	if leftIndex >= rightIndex {
		return
	}

	half := (rightIndex + leftIndex) / 2
	MergeSortLowMemVersion(nums, leftIndex, half)
	MergeSortLowMemVersion(nums, half+1, rightIndex)

	for i := half + 1; i <= rightIndex; i++ {
		temp := nums[i]
		j := i - 1
		for ; j >= leftIndex; j-- {
			if nums[j] <= temp {
				break
			} else {
				nums[j+1] = nums[j]
			}
		}
		nums[j+1] = temp
	}
	fmt.Println(nums)
}
