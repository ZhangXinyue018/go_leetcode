package multiply_strings

import "strings"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	temp := make([]int, l1+l2+1)
	for j := l1 - 1; j >= 0; j-- {
		for i := l2 - 1; i >= 0; i-- {
			baseIndex := l1 + l2 - 2 - i - j
			num := int(num2[i]-'0') * int(num1[j]-'0')
			AddToTemp(&temp, baseIndex, num)
		}
	}
	sb := strings.Builder{}
	i := len(temp) - 1
	for ; i >= 0; i-- {
		if temp[i] != 0 {
			break
		}
	}
	for ; i >= 0; i-- {
		sb.WriteByte(byte('0' + temp[i]))
	}
	return sb.String()
}

func AddToTemp(temp *[]int, baseIndex int, num int) {
	baseNumber := (*temp)[baseIndex]
	baseNumber = baseNumber + num
	nextNumber := (*temp)[baseIndex+1]
	nextNumber = nextNumber + baseNumber/10
	(*temp)[baseIndex] = baseNumber % 10
	(*temp)[baseIndex+1] = nextNumber
}
