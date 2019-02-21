package panlindrome_number

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	var result int64 = 0
	var input = int64(x)
	for input > result {
		result = result*10 + input%10
		input = input / 10
	}
	return input == result || input == result/10
}
