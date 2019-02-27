package permutation

func Permutation(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	result := []string{}
	for i := 0; i < len(s); i++ {
		newPreAppend := string(s[i])
		tempResults := Permutation(s[:i]+s[i+1:])
		for _, tempResult := range tempResults {
			result = append(result, newPreAppend + tempResult)
		}
	}
	return result
}
