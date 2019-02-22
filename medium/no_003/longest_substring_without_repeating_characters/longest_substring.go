package longest_substring_without_repeating_characters

// ************* Problem #3 *************

func lengthOfLongestSubstring(s string) int {
	lastRepeat := -1
	longestLen := 0
	indexMap := map[int32]int{}
	for i, char := range s {
		if index, ok := indexMap[char]; ok && lastRepeat < index {
			lastRepeat = index
		}
		indexMap[char] = i
		if length := i - lastRepeat; length > longestLen {
			longestLen = length
		}
	}
	return longestLen
}

func lengthOfLongestSubstring2(s string) int {
	countIndex := -1
	maxLength := 0
	markingMap := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if index, ok := markingMap[s[i]]; ok {
			if index > countIndex {
				countIndex = index
			}
		}
		markingMap[s[i]] = i
		tempLength := i - countIndex
		if tempLength > maxLength {
			maxLength = tempLength
		}
	}
	return maxLength
}
