package main

// ************* Problem #3 *************

import "fmt"

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

func main() () {
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("ababab"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
