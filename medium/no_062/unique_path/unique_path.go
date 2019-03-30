package unique_path

func uniquePaths(m int, n int) int {
	mMap := make([]int, m)
	DetailPath(m, n, &mMap)
	return mMap[len(mMap)-1]
}

func DetailPath(m int, n int, resultMap *[]int) {
	for j := 1; j <= n; j ++ {
		if j == 1 {
			for i := 0; i < len(*resultMap); i++ {
				(*resultMap)[i] = 1
			}
		} else {
			for i := 0; i < len(*resultMap); i++ {
				if i != 0 {
					(*resultMap)[i] = (*resultMap)[i-1] + (*resultMap)[i]
				}
			}

		}
	}
}
