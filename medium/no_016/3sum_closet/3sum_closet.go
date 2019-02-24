package _sum_closet

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	closeSum := 0
	closeDiff := math.MaxInt32
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			//if left > i+1 && nums[left] == nums[left-1] {
			//	left++
			//	continue
			//}
			//if right < len(nums)-1 && nums[right] == nums[right+1] {
			//	right--
			//	continue
			//}
			tempSum := nums[i] + nums[left] + nums[right]
			if tempSum == target {
				return target
			} else if tempSum > target {
				right --
				if tempSum-target < closeDiff {
					closeDiff = tempSum - target
					closeSum = tempSum
				}
			} else if tempSum < target {
				left ++
				if target-tempSum < closeDiff {
					closeDiff = target - tempSum
					closeSum = tempSum
				}
			}
		}
	}
	return closeSum
}
