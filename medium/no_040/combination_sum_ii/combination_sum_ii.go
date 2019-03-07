package combination_sum_ii

import "sort"

// todo: implement 

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	DetailedCombination(candidates, []int{}, target, 0, &result)
	return result
}

func DetailedCombination(candidates []int, output []int, target int, index int, result *[][]int) {
	if target == 0 {
		*result = append(*result, output)
		return
	} else if target < 0 || index >= len(candidates) {
		return
	}

	DetailedCombination(candidates, append([]int{candidates[index]}, output...), target-candidates[index], index+1, result)
	i := index + 1
	for ; i < len(candidates); i++ {
		if candidates[i] != candidates[index] {
			break
		}
	}
	DetailedCombination(candidates, output, target, i, result)

}
