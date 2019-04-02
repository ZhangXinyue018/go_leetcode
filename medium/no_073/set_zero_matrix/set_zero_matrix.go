package set_zero_matrix

func setZeroes(matrix [][]int) {
	fr, fc := false, false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if j != 0 && i != 0 {
					matrix[i][0] = 0
					matrix[0][j] = 0
					continue
				}
				if j == 0 {
					fc = true
				}
				if i == 0 {
					fr = true
				}
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if fr {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if fc {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
