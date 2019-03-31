package min_path_sum

func minPathSum(grid [][]int) int {
	result := make([]int, len(grid))
	for j := len(grid[0]) - 1; j >= 0; j-- {
		for i := len(grid) - 1; i >= 0; i-- {
			if i == len(grid)-1 {
				result[i] = grid[i][j] + result[i]
			} else if j == len(grid[0])-1 {
				result[i] = result[i+1] + grid[i][j]
			} else {
				result[i] = min(result[i], result[i+1]) + grid[i][j]
			}
		}
	}
	return result[0]
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
