package combination_sum

import "sort"

// - this method is to form a record map
// - each key of the map represents a sum value
// - recurse candidates (candidates may recurse many times)
// - find key with target of map
// I was trying to mark down answer of each sum so that there won't be duplicate calculations
// However, it turns out that too many times of copying slices will cause in even lower speed
// This method is not recommended
func combinationSum(candidates []int, target int) [][]int {
	combinationMap := map[int][][]int{}
	for _, candidate := range candidates {
		if _, hasKey := combinationMap[candidate]; hasKey {
			combinationMap[candidate] = append(combinationMap[candidate], []int{candidate})
		} else {
			combinationMap[candidate] = [][]int{{candidate}}
		}

		for i := 0; i <= target; i++ {
			value, hasResult := combinationMap[i]
			if hasResult {
				for _, temp := range value {
					if _, hasKey := combinationMap[i+candidate]; hasKey {
						combinationMap[i+candidate] = append(
							[][]int{append([]int{candidate}, temp...)},
							combinationMap[i+candidate]...
						)
					} else {
						combinationMap[i+candidate] = [][]int{append([]int{candidate}, temp...)}
					}
				}
			}
		}
	}
	return combinationMap[target]
}

// - to find combination of sum is to find
// 		(combination of sum - firstnum) append firstnum AND
//	 	combination of sum (without firstnum)
// In this question firstnum can be included many times
// This method is clearer and faster in speed
// Recommend
func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	DetailCombination2(candidates, []int{}, 0, target, &result)
	return result
}

func DetailCombination2(candidates []int, output []int, index int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, output...))
		return
	} else if index == len(candidates) || target < 0 {
		return
	}

	if target >= candidates[index] {
		DetailCombination2(candidates, append(output, candidates[index]), index, target-candidates[index], result)
	}
	DetailCombination2(candidates, output, index+1, target, result)
}

// Similar as previous one
// - To find combination of sum is to find
//			every:
//				(combination of sum-index) append num(index)
// This method is clearer
// I tried to sort candidates before finding combination
// It turns out no difference on memory usage TAT
// But method 2 & method 3 faster than 100% of Go submissions
// Recommend
func combinationSum3(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	DetailCombination3(candidates, []int{}, 0, target, &result)
	return result
}

func DetailCombination3(candidates []int, output []int, index int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, output...))
		return
	}

	for i := index; i < len(candidates) && candidates[i] <= target; i++ {
		DetailCombination3(candidates, append(output, candidates[i]), i, target-candidates[i], result)
	}

}
