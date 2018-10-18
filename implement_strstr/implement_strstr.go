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

