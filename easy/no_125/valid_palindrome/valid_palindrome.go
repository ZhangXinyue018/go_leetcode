package valid_palindrome

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	leftPointer := 0
	rightPointer := len(s) - 1
	for leftPointer < rightPointer {
		if !IsAlphaNumeric(s[leftPointer]) {
			leftPointer++
			continue
		} else if !IsAlphaNumeric(s[rightPointer]) {
			rightPointer--
			continue
		}
		if s[leftPointer] != s[rightPointer] {
			return false
		} else {
			leftPointer++
			rightPointer--
		}
	}
	return true
}

func IsAlphaNumeric(s uint8) bool {
	return (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9')
}
