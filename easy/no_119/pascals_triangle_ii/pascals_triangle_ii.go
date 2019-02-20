package pascals_triangle_ii

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	resultRow := getRow(rowIndex - 1)
	resultRow = append(resultRow, 1)
	for i := rowIndex - 1; i > 0; i-- {
		resultRow[i] = resultRow[i] + resultRow[i-1]
	}
	return resultRow
}
