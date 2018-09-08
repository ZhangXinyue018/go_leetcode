package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	prefix := strs[0]
	length := len(strs)
	for i := 1; i < length; i++ {
		str := strs[i]
		strLenth := len(str)
		prefixLength := len(prefix)
		temp := []uint8{}
		for j := 0; j < prefixLength && j < strLenth; j++ {
			if str[j] == prefix[j] {
				temp = append(temp, str[j])
			}else{
				break
			}
		}
		prefix = string(temp)
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	prefix := strs[0]
	length := len(strs)
	for i := 1; i < length; i++ {
		str := strs[i]
		for !strings.HasPrefix(str, prefix){
			if len(prefix) == 1 {
				prefix = ""
				break
			}
			prefix = prefix[0:len(prefix)-1]
		}
	}
	return prefix
}


func main() () {
	fmt.Println(longestCommonPrefix2([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix2([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix2([]string{}))
	fmt.Println(longestCommonPrefix2([]string{"aca", "cba"}))
	fmt.Println(longestCommonPrefix2([]string{"aa", "ab"}))
}
