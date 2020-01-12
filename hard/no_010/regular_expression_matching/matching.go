package regular_expression_matching

func isMatch(s string, p string) bool {
	// less ram, more time
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(s) == 0 || (p[0] != '.' && s[0] != p[0]) {
		if len(p) >= 2 && p[1] == '*' {
			return isMatch(s, p[2:])
		} else {
			return false
		}
	} else {
		if len(p) >= 2 && p[1] == '*' {
			return isMatch(s[1:], p) || isMatch(s, p[2:])
		} else {
			return isMatch(s[1:], p[1:])
		}
	}
}

func isMatch2(s string, p string) bool {
	// less time, more ram
	resultTable := [][]bool{}
	for i := 0; i <= len(s); i++ {
		resultTable = append(resultTable, make([]bool, len(p)+1))
	}
	resultTable[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[len(p)-j] == '*' {
			j++
			for i := 0; i <= len(s); i++ {
				resultTable[i][j] = resultTable[i][j-2]
				if i > 0 && (p[len(p)-j] == '.' || s[len(s)-i] == p[len(p)-j]) {
					resultTable[i][j] = resultTable[i][j] || resultTable[i-1][j]
				}
			}
		} else {
			has := false
			for i := 1; i <= len(s); i++ {
				if p[len(p)-j] == '.' || s[len(s)-i] == p[len(p)-j] {
					has = true
					resultTable[i][j] = resultTable[i-1][j-1]
				}
			}
			if !has {
				return false
			}
		}
	}
	return resultTable[len(s)][len(p)]
}
