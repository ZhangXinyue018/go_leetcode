package unique_path_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	tempMap := make([]int, len(obstacleGrid))
	for j := len(obstacleGrid[0]) - 1; j >= 0; j-- {
		for i := len(obstacleGrid) - 1; i >= 0; i-- {
			if obstacleGrid[i][j] == 1 {
				tempMap[i] = 0
			} else if j == len(obstacleGrid[0])-1 {
				if i == len(obstacleGrid)-1 || tempMap[i+1] == 1 {
					tempMap[i] = 1
				}
			} else if i != len(obstacleGrid)-1 {
				tempMap[i] = tempMap[i] + tempMap[i+1]
			}
		}
	}
	return tempMap[0]
}
