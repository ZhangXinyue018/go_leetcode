package two_sum_ii

func twoSum(numbers []int, target int) []int {
	temp := make(map[int]int)
	for index, value := range numbers {
		if lastIndex, ok := temp[value]; ok {
			return []int{lastIndex + 1, index + 1}
		} else {
			temp[target-value] = index
		}
	}
	return []int{0, 0}
}
