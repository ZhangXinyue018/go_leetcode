package wildcard_matching

func isMatch(s string, p string) bool {
	holdingIndexes := []int{len(s)}
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] == '*' {
			max := holdingIndexes[len(holdingIndexes)-1]
			newIndex := []int{}
			for j := 0; j <= max; j++ {
				newIndex = append(newIndex, j)
			}
			holdingIndexes = newIndex
		} else {
			for j := 0; j < len(holdingIndexes); j++ {
				index := holdingIndexes[j]
				if index != 0 && (p[i] == '?' || s[index-1] == p[i]) {
					holdingIndexes[j] = holdingIndexes[j] - 1
				} else {
					holdingIndexes = append(holdingIndexes[:j], holdingIndexes[j+1:]...)
					j--
				}
			}
		}
		if len(holdingIndexes) == 0 {
			return false
		}
	}
	return len(holdingIndexes) > 0 && holdingIndexes[0] == 0
}
