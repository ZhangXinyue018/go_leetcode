package zigzag_conversion

import (
	"strings"
)

func Convert(s string, numRows int) string {
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

func ConvertBetter(s string, numRows int) string {
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

func ConvertTest(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var sb strings.Builder
	flipNumber := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		if i != 0 && i != numRows-1 {
			swing := 2 * i
			for j := i; j < len(s); j = j + swing {
				sb.WriteString(string(s[j]))
				swing = flipNumber - swing
			}
		} else {
			for j := i; j < len(s); j = j + flipNumber {
				sb.WriteString(string(s[j]))
			}
		}
	}
	return sb.String()
}
