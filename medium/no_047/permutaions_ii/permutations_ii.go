package permutaions_ii

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	preNumber := 0
	finalResult := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == preNumber {
			continue
		}
		temp := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		prePermuteResults := permuteUnique(temp)
		for _, prePermuteResult := range prePermuteResults {
			finalResult = append(finalResult, append([]int{nums[i]}, prePermuteResult...))
		}
		preNumber = nums[i]
	}
	return finalResult

}

func permuteUnique3(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	recordMap := map[int]bool{}
	finalResult := [][]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := recordMap[nums[i]]; ok {
			continue
		}
		nums[0], nums[i] = nums[i], nums[0]
		prePermuteResults := permuteUnique3(nums[1:])
		for _, prePermuteResult := range prePermuteResults {
			finalResult = append(finalResult, append([]int{nums[0]}, prePermuteResult...))
		}
		nums[0], nums[i] = nums[i], nums[0]
		recordMap[nums[i]] = true
	}
	return finalResult

}

// incorrect
func permuteUnique2(nums []int) [][]int {
	//if len(nums) == 1 {
	//	return [][]int{{nums[0]}}
	//}
	//
	//preResults := permuteUnique2(nums[1:])
	//finalResult := [][]int{}
	//for _, preResult := range preResults {
	//	finalResult = append(finalResult, append([]int{nums[0]}, preResult...))
	//	for j := 1; j <= len(preResult); j++ {
	//		if preResult[j-1] == nums[0] {
	//			continue
	//		}
	//		finalResult = append(finalResult, append(append(append([]int{}, preResult[:j]...), nums[0]), preResult[j:]...))
	//	}
	//}
	//return finalResult
}
