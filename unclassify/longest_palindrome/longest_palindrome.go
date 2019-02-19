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
