package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	resultMap := map[int][]string{}
	returnPoint := numRows - 1
	countInGroup := returnPoint << 1
	if numRows == 1 {
		return s
	}
	subPointer := 0
	for _, char := range s {
		if subPointer == countInGroup {
			subPointer = 0
		}
		if subPointer <= returnPoint {
			resultMap[subPointer] = append(resultMap[subPointer], string(char))
		} else {
			resultMap[countInGroup-subPointer] = append(resultMap[countInGroup-subPointer], string(char))
		}
		subPointer++
	}
	resultString := ""
	for i := 0; i < numRows; i++ {
		for _, str := range resultMap[i] {
			resultString = resultString + str
		}
	}
	return resultString
}

func convertBetter(s string, numRows int) string {
	var sb strings.Builder
	if numRows == 1 {
		return s
	}
	strLen := len(s)
	for i := 0; i < numRows; i++ {
		j := i
		if j == 0 || j == numRows-1 {
			for j < strLen {
				sb.WriteByte(byte(s[j]))
				j = j + 2*numRows - 2
			}
		} else {
			swing := numRows - 1 - i
			for j < strLen {
				sb.WriteByte(byte(s[j]))
				j = j + 2*swing
				swing = numRows - 1 - swing
			}
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(convertBetter("PAYPALISHIRING", 3))
	fmt.Println(convertBetter("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convertBetter("PAYPALISHIRING", 4))
	fmt.Println(convertBetter("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
	fmt.Println(convertBetter("AB", 1))
	fmt.Println(convertBetter("AB", 1) == "AB")
}
