package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxValue := nums[0]
	var maxCurrent int
	if nums[0] > 0{
		maxCurrent = nums[0]
	}else{
		maxCurrent = 0
	}
	numLen := len(nums)
	for i := 1; i < numLen; i++ {
		maxCurrent = maxCurrent + nums[i]
		if maxValue < maxCurrent {
			maxValue = maxCurrent
		}
		if maxCurrent < 0 {
			maxCurrent = 0
		}
	}
	return maxValue
}

const INT_MIN = math.MinInt32

func maxSubArray2(nums []int) int {
	max := INT_MIN
	sum := 0
	for _, value := range nums {
		sum = sum + value
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1}))
}
