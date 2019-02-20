package pascals_triangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := [][]int{[]int{1}}
	for i := 1; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j <= i/2; j++ {
			temp[j] = result[i-1][j-1] + result[i-1][j]
			temp[i-j] = temp[j]
		}
		result = append(result, temp[:])
	}
	return result
}
