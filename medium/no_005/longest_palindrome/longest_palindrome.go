package longest_palindrome

func longestPalindrome(s string) string {
	strLen := len(s)
	longestLength := 0
	longestStr := ""
	for i := 0; i < strLen; i++ {
		currLength := 0
		j, k := i, i
		for k < strLen-1 {
			if s[k+1] != s[i] {
				break
			}
			k++
		}
		for j > 0 && k < strLen-1 {
			if s[j-1] != s[k+1] {
				break
			}
			j--
			k++
		}
		currLength = k - j + 1
		if currLength > longestLength {
			longestLength = currLength
			longestStr = s[j : k+1]
		}
	}
	return longestStr
}

func longestPalindromeBetter(s string) string {
	resultStr := ""
	for i := 0; i < len(s); i++ {
		j := i
		for ; j < len(s)-1; j++ {
			if s[i] != s[j+1] {
				break
			}
		}
		tempStr := getLongest(s, i, j)
		if len(tempStr) > len(resultStr) {
			resultStr = tempStr
		}
	}
	return resultStr
}

func getLongest(s string, leftStartPointer int, rightStartPointer int) string {
	for leftStartPointer > 0 && rightStartPointer < len(s)-1 {
		if s[leftStartPointer-1] != s[rightStartPointer+1] {
			break
		}
		leftStartPointer--
		rightStartPointer++
	}
	return s[leftStartPointer : rightStartPointer+1]
}
