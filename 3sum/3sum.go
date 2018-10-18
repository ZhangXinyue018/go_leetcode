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
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else {
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
			}
		}
	}
	return resultList
}
