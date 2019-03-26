package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0{
		return []int{}
	}
	m := len(matrix)-1
	n := len(matrix[0])-1

	result := []int{}
	i,j := 0, -1
	count :=0
	jflip := 1
	iflip := 1
	for m >=0 && n >=0{
		if count %2 ==0{
			for temp := 0; temp <= n;temp++{
				j = j + jflip
				result = append(result, matrix[i][j])
			}
			m--
			jflip = -jflip
		}else{
			for temp := 0; temp <= m; temp++{
				i = i + iflip
				result = append(result, matrix[i][j])
			}
			n--
			iflip = -iflip
		}
		count++
		if count == 4{
			count = 0
		}
	}
	return result
}
