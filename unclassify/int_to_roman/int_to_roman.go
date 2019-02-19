package int_to_roman

import (
	"strings"
	)

var intList = []int{1000, 500, 100, 50, 10, 5, 1}
var romanList = []uint32{'M', 'D', 'C', 'L', 'X', 'V', 'I'}

func intToRoman(num int) string {
	var sb strings.Builder
	length := len(intList)
	i := 0
	for i < length {
		for num >= intList[i] {
			sb.WriteByte(byte(romanList[i]))
			num = num - intList[i]
		}
		j := i + 1
		for j < length {
			if num >= (intList[i]-intList[j]) && (intList[i]-intList[j]) != intList[j] {
				sb.WriteByte(byte(romanList[j]))
				sb.WriteByte(byte(romanList[i]))
				num = num - intList[i] + intList[j]
			}
			j++
		}
		i++
	}
	for num > 1 {
		sb.WriteByte(byte('I'))
		num = num - 1
	}
	return sb.String()
}
