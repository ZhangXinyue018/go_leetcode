package main

import (
	"fmt"
	"strings"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	resultStr := "1"
	for i := 2; i <= n; i++ {
		resultStr = returnValue(resultStr)
	}
	return resultStr
}

func returnValue(value string) string {
	var sb strings.Builder
	lastChar, lastCount := value[0], 1
	strLen := len(value)
	for i := 1; i < strLen; i++ {
		char := value[i]
		if char == lastChar {
			lastCount ++
		} else {
			sb.WriteString(strconv.Itoa(lastCount))
			sb.WriteByte(byte(lastChar))
			lastChar = char
			lastCount = 1
		}
	}
	sb.WriteString(strconv.Itoa(lastCount))
	sb.WriteByte(byte(lastChar))
	return sb.String()
}

func main() {
	for i := 1; i <= 30; i ++ {
		fmt.Printf("%d: %s\n", i, countAndSay(i))
	}

}
