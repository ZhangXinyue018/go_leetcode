package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	row := [9][9]byte{}
	column := [9][9]byte{}
	subBox := [9][9]byte{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row[i][j] = board[i][j]
			column[i][j] = board[j][i]
			subBox[int(j/3)*3+(i/3)][int(i%3)*3+j%3] = board[i][j]
		}
	}
	for i := 0; i < 9; i++ {
		if !isBoxValid(row[i][:]) || !isBoxValid(column[i][:]) || !isBoxValid(subBox[i][:]) {
			return false
		}
	}
	return true
}

func isBoxValid(box []byte) bool {
	tempMap := map[byte]int{}
	for _, boxNum := range box {
		if boxNum == '.' {
			continue
		}
		if _, hasValue := tempMap[boxNum]; hasValue {
			return false
		} else {
			tempMap[boxNum] = 1
		}
	}
	return true
}
