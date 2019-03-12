package longest_common_string

func LongestCommonSubSequence(a string, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	if a[len(a)-1] == b[len(b)-1] {
		return LongestCommonSubSequence(a[:len(a)-1], b[:len(b)-1]) + 1
	} else {
		result1 := LongestCommonSubSequence(a[:len(a)-1], b)
		result2 := LongestCommonSubSequence(a, b[:len(b)-1])
		if result1 > result2{
			return result1
		}else{
			return result2
		}
	}
}
