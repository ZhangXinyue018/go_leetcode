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
