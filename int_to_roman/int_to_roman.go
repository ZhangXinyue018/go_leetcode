package main

import (
	"strings"
	"fmt"
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

func main() {
	fmt.Println(intToRoman(8))    //VIII
	fmt.Println(intToRoman(3))    //III
	fmt.Println(intToRoman(4))    //IV
	fmt.Println(intToRoman(9))    //IX
	fmt.Println(intToRoman(58))   //LVIII
	fmt.Println(intToRoman(1994)) //MCMXCIV
}
