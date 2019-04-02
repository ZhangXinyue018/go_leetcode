package word_search

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	usedMap := [][]bool{}
	for i := 0; i < m; i++ {
		usedMap = append(usedMap, make([]bool, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == byte(word[0]) {
				usedMap[i][j] = true
				result := dfs(board, m, n, i, j, 1, word, usedMap)
				if result {
					return true
				}
				usedMap[i][j] = false
			}
		}
	}
	return false
}

func dfs(board [][]byte, m, n, i, j, pos int, word string, usedMap [][]bool) bool {
	if pos == len(word) {
		return true
	}
	type point struct {
		X int
		Y int
	}
	tasks := []point{}
	if i > 0 && board[i-1][j] == byte(word[pos]) && !usedMap[i-1][j] {
		tasks = append(tasks, point{i - 1, j})
	}
	if i < m-1 && board[i+1][j] == byte(word[pos]) && !usedMap[i+1][j] {
		tasks = append(tasks, point{i + 1, j})
	}
	if j > 0 && board[i][j-1] == byte(word[pos]) && !usedMap[i][j-1] {
		tasks = append(tasks, point{i, j - 1})
	}
	if j < n-1 && board[i][j+1] == byte(word[pos]) && !usedMap[i][j+1] {
		tasks = append(tasks, point{i, j + 1})
	}
	for _, task := range tasks {
		usedMap[task.X][task.Y] = true
		result := dfs(board, m, n, task.X, task.Y, pos+1, word, usedMap)
		if result {
			return true
		}
		usedMap[task.X][task.Y] = false
	}
	return false
}
