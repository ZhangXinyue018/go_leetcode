package combination

func combine(n int, k int) [][]int {
	result := [][]int{}
	tempList := []Pair{{[]int{}, 0}}
	for i := n; i >= 1; i-- {
		resultTempList := []Pair{}
		for _, temp := range tempList {
			if i+temp.K < k {
				continue
			}
			valueAfterAdd := append([]int{i}, temp.Val...)
			kAfterAdd := temp.K + 1
			if kAfterAdd == k {
				result = append(result, valueAfterAdd)
			} else {
				resultTempList = append(resultTempList, Pair{valueAfterAdd, kAfterAdd})
			}
			resultTempList = append(resultTempList, temp)
		}
		tempList = resultTempList
	}
	return result
}

type Pair struct {
	Val []int
	K   int
}
