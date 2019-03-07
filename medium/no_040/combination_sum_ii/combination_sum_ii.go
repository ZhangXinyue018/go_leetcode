package combination_sum_ii

import "sort"

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

func combinationSum22(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return DetailedCombination22(candidates, target, 0)
}

func DetailedCombination22(candicates []int, target int, index int) [][]int {
	if target == 0 {
		return [][]int{{}}
	} else if target < 0 || index >= len(candicates) {
		return [][]int{}
	}

	finalresult := [][]int{}
	for i := index; i < len(candicates); i++ {
		if i > index && candicates[i] == candicates[i-1] {
			continue
		} else if candicates[i] > target {
			break
		} else {
			tempResult := DetailedCombination22(candicates, target-candicates[i], i+1)
			for _, temp := range tempResult {
				finalresult = append(finalresult, append(temp, candicates[i]))
			}
		}
	}
	return finalresult
}
