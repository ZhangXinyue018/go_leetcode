package searc_2d_matrix

import (
	"sync"
)

type Result struct {
	Val bool
}

func searchMatrix2(matrix [][]int, target int) bool {
	result := &Result{}
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	var wg sync.WaitGroup
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			wg.Add(1)
			go func(matrix [][]int, i, j, target int, result *Result) {
				defer wg.Done()
				if matrix[i][j] == target {
					(*result).Val = true
				}
			}(matrix, i, j, target, result)
		}
	}
	wg.Wait()
	return result.Val
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	find := make(chan bool, 1)

	go func(matrix [][]int, m, n, target int, find chan bool) {
		wg := &sync.WaitGroup{}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				wg.Add(1)
				go Find(matrix, i, j, target, find, wg)
			}
		}
		wg.Wait()
		close(find)
	}(matrix, m, n, target, find)

	for a := range find {
		return a
	}
	return false

}

func Find(matrix [][]int, i, j, target int, find chan bool, wg *sync.WaitGroup) {
	if matrix[i][j] == target {
		find <- true
	}
	wg.Done()
}


func searchMatrix3(matrix [][]int, target int) bool {
	if len(matrix) == 0{
		return false
	}
	startR := 0
	endR := len(matrix) -1
	for startR < endR{
		mid := (startR+endR+1)/2
		if matrix[mid][0] == target{
			return true
		}else if matrix[mid][0] > target{
			endR = mid -1
		}else{
			startR = mid
		}
	}
	startC := 0
	endC := len(matrix[0])-1
	for startC <= endC{
		mid := (startC+endC)/2
		if matrix[startR][mid] == target{
			return true
		}else if matrix[startR][mid] > target{
			endC = mid-1
		}else{
			startC = mid+1
		}
	}
	return false
}