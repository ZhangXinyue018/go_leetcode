package main

import (
	"strings"
	)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	} else if len(haystack) < len(needle) {
		return -1
	}
	length := len(haystack) - len(needle)
	for i := 0; i <= length; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func strStr3(haystack string, needle string) int{
	if needle == "" {
		return 0
	}
	haystackrune := []rune(haystack)
	needlerune := []rune(needle)
	for i := 0; i <= (len(haystackrune) - len(needlerune)); i++{
		for j := 0; j < len(needlerune); j++{
			if haystackrune[i+j] != needlerune[j]{
				break
			}else if j == len(needlerune) -1{
				return i
			}
		}
	}
	return -1
}
