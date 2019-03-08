package permutations

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	preResults := permute(nums[1:])
	var finalResult [][]int
	for _, preResult := range preResults {
		finalResult = append(finalResult, append([]int{nums[0]}, preResult...))
		for j := 1; j <= len(preResult); j++ {
			finalResult = append(finalResult, append(append(append([]int{}, preResult[:j]...), nums[0]), preResult[j:]...))
		}
	}
	return finalResult
}
