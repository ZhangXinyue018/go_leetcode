package factorial_trailing_zeroes

func trailingZeroes(n int) int {
	result := 0
	for n != 0 {
		result += n / 5
		n = n / 5
	}
	return result
}
