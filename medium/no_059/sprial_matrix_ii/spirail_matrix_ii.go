package sprial_matrix_ii

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for temp := 0;temp < len(result); temp++{
		result[temp] = make([]int, n)
	}
	iMax, jMax := n, n
	i, j := 0, -1
	iFlip, jFlip := 1, 1
	count := 1
	num := 1
	for iMax > 0 && jMax > 0{
		if count %2 == 1{
			for temp := 0; temp < jMax;temp++{
				j = j + jFlip
				result[i][j] = num
				num++
			}
			jFlip = -jFlip
			iMax--
		}else{
			for temp := 0; temp < iMax;temp++{
				i = i + iFlip
				result[i][j] = num
				num++
			}
			iFlip = -iFlip
			jMax--
		}
		count++
	}
	return result
}
