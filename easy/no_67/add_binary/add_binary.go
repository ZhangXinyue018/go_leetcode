package add_binary

import (
	"strconv"
)

var (
	CharMap = map[rune]int{
		'1': 1,
		'0': 0,
	}
)

var (
	UniCharMap = map[uint8]int{
		'1': 1,
		'0': 0,
	}
)

func addBinary(a string, b string) string {
	tempA := []rune(a)
	tempB := []rune(b)
	i := len(tempA) - 1
	j := len(tempB) - 1
	cin := 0
	var resultStr []rune
	for i >= 0 || j >= 0 || cin > 0 {
		valueA := 0
		valueB := 0
		if i >= 0 {
			valueA = CharMap[tempA[i]]
			i--
		}
		if j >= 0 {
			valueB = CharMap[tempB[j]]
			j--
		}
		temp := valueA + valueB + cin
		cin = temp / 2
		resultStr = append([]rune(strconv.Itoa(temp%2)), resultStr...)
	}
	return string(resultStr)
}

func addBinary2(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	cin := byte(0)
	resultStr := ""
	for i >= 0 || j >= 0 || cin > 0 {
		sum := cin
		if i >= 0 {
			sum = sum + a[i] - '0'
			i--
		}
		if j >= 0 {
			sum = sum + b[j] - '0'
			j--
		}
		cin = sum / 2

		resultStr = string('0'+sum%2) + resultStr
	}
	return resultStr
}
