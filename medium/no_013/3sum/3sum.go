package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var resultList [][]int
	sort.Ints(nums)
	numLen := len(nums)
	for i := 0; i < numLen-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i != 0 && nums[i] == nums [i-1] {
			continue
		}
		left, right := i+1, numLen-1
		for left < right {
			temp := nums[i] + nums[left] + nums[right]
			if temp == 0 {
				resultList = append(resultList, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if temp > 0 {
				//for left < right && nums[right] == nums[right-1] {
				//	right--
				//}
				right--
			} else {
				//for left < right && nums[left] == nums[left+1] {
				//	left++
				//}
				left++
			}
		}
	}
	return resultList
}

// todo: this answer give duplicate result, where the duplicate results are from difference indexes
// todo: not recommended
func threeSumTest(nums []int) [][]int {
	return findSum(nums, []int{}, 0, len(nums)-1, 0, 3)
}

func findSum(nums []int, preList []int, fromIndex int, toIndex int, sum int, count int) [][]int {
	if count < 0 {
		return [][]int{}
	} else if count == 0 {
		if sum == 0 {
			return [][]int{preList}
		} else {
			return [][]int{}
		}
	} else if count == 1 {
		if fromIndex == toIndex && sum == nums[fromIndex] {
			return [][]int{append(preList, nums[fromIndex])}
		}
	}
	if fromIndex > toIndex {
		return [][]int{}
	}

	return append(findSum(nums, append(preList, nums[fromIndex]), fromIndex+1, toIndex, sum-nums[fromIndex], count-1),
		findSum(nums, preList, fromIndex+1, toIndex, sum, count)...)
}
