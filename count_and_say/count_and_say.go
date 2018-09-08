package main

import "fmt"

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
	resultString := ""
	lastChar := value[0]
	lastCount := 1
	strLen := len(value)
	for i := 1; i < strLen; i++ {
		char := value[i]
		if char == lastChar {
			lastCount ++
		} else {
			resultString = resultString + string('0'+lastCount) + string(lastChar)
			lastChar = char
			lastCount = 1
		}
	}
	return resultString + string('0'+lastCount) + string(lastChar)
}

func main() {
	for i := 1; i <= 30; i ++ {
		fmt.Printf("%d: %s\n", i, countAndSay(i))
	}

}
