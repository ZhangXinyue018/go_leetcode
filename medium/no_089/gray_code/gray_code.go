package gray_code

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	result := []int{0, 1}
	pw := 2
	for i := 2; i <= n; i++ {
		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, result[j]+pw)
		}
		pw = pw << 1
	}
	return result
}
