package decorate_ways

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	pre, curr := 1, 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			pre = curr
			curr = 0
			continue
		} else if i == len(s)-1 {
			continue
		}
		tempStr := string(s[i : i+2])
		if tempStr == "20" || tempStr == "10" {
			tempCurr := pre
			pre = curr
			curr = tempCurr
		} else if tempStr > "26" {
			pre = curr
		} else {
			tempCurr := curr + pre
			pre = curr
			curr = tempCurr
		}
	}
	return curr
}
