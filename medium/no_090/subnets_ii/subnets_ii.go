package subnets_ii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return DetailSubSet(nums)
}

func DetailSubSet(nums []int) [][]int {
	result := [][]int{[]int{}}
	temp := [][]int{[]int{nums[0]}}
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			temp = append(temp, append([]int{pre}, temp[len(temp)-1]...))
		} else {
			formatResult(&result, temp)
			temp = [][]int{[]int{nums[i]}}
		}
		pre = nums[i]
	}
	formatResult(&result, temp)
	return result
}

func formatResult(result *[][]int, temp [][]int) {
	resultLen := len(*result)
	for j := 0; j < resultLen; j++ {
		for _, tempV := range temp {
			appendSlice := append([]int{}, (*result)[j]...)
			appendSlice = append(appendSlice, tempV...)
			*result = append(*result, appendSlice)
		}
	}
}
