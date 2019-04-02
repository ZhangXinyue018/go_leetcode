package word_search

func exist_ii(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++{
		for j :=0;j<n; j++{
			result := dfs_ii(board, m,n,i,j, 0, word)
			if result{
				return true
			}
		}
	}
	return false
}

func dfs_ii(board [][]byte, m, n, i, j, pos int, word string) bool{
	if pos >= len(word) || i < 0 || i >= m || j < 0 || j >=n{
		return false
	}else if board[i][j] != byte(word[pos]){
		return false
	}else if pos == len(word) -1{
		return true
	}
	type point struct{
		X int
		Y int
	}
	tasks := []point{
		{i-1, j},
		{i+1, j},
		{i, j-1},
		{i, j+1},
	}
	board[i][j] ^= 255
	result := false
	for _, task := range tasks{
		result = dfs_ii(board, m, n, task.X, task.Y, pos+1, word)
		if result{
			break
		}
	}
	board[i][j] ^= 255
	return result
}
