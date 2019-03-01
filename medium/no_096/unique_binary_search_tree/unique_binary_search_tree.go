package unique_binary_search_tree

import "strconv"

var (
	composeKey = make(map[string]int)
)

func numTrees(n int) int {
	return DetailNumTrees(1, n)
}

func DetailNumTrees(fromIndex int, toIndex int) int {
	key := strconv.Itoa(fromIndex) + "-" + strconv.Itoa(toIndex)
	if value, ok := composeKey[key]; ok {
		return value
	}
	result := 0
	if fromIndex >= toIndex {
		result = 1
	} else {
		for i := fromIndex; i <= toIndex; i++ {
			result = result + DetailNumTrees(fromIndex, i-1)*DetailNumTrees(i+1, toIndex)
		}
	}
	composeKey[key] = result
	return result

}
