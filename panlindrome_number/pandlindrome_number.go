package main

import "fmt"

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

func main() () {
	fmt.Println(isPalindrome(-123))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(10))
}
