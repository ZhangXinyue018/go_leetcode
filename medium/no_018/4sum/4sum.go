package _sum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return detailSum(nums, target, 4)
}

func detailSum(nums []int, target int, count int) [][]int {
	if len(nums) == 0 || count < 1 {
		return [][]int{}
	}
	if (nums[0] > 0 && target < nums[0]) ||
		(nums[len(nums)-1] <0 && target > nums[len(nums)-1]){
		return [][]int{}
	}
	//if count == 1 {
	//	for _, num := range nums {
	//		if num == target {
	//			return [][]int{{num}}
	//		}
	//	}
	//	return [][]int{}
	//}
	if count == 2 {
		result := [][]int{}
		left := 0
		right := len(nums) - 1
		for left < right {
			if left > 0 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			sum := nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum == target {
				result = append(result, []int{nums[left], nums[right]})
				left++
				right--
			} else {
				left++
			}
		}
		return result
	}
	j := 1
	for ; j < len(nums); j++ {
		if nums[j] != nums[0] {
			break
		}
	}
	resultsContain := detailSum(nums[1:], target-nums[0], count-1)
	result := [][]int{}
	for _, resultContain := range resultsContain {
		result = append(result, append(resultContain, nums[0]))
	}
	if j < len(nums) {
		resultsNotContain := detailSum(nums[j:], target, count)
		result = append(result, resultsNotContain...)
	}
	return result
}
