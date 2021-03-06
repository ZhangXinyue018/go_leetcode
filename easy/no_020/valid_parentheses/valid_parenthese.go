package valid_parentheses

var (
	Matcher = map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}
)

func IsValid(s string) bool {
	queue := []int32{}
	for _, char := range s {
		if result, hasRecord := Matcher[char]; hasRecord {
			if len(queue) > 0 && queue[len(queue)-1] == result {
				queue = queue[0 : len(queue)-1]
			} else {
				return false
			}
		} else {
			queue = append(queue, char)
		}
	}
	return len(queue) == 0
}

