package main

import (
	"fmt"
	"math"
)

const (
	StageSign  = 1
	StageValue = 2
)

func myAtoi(str string) int {
	stage := StageSign
	sign := 1
	var resultInt int64 = 0
	for _, char := range str {
		if stage == StageSign {
			if char == ' ' {
				continue
			} else if char == '+' {
				sign = 1
				stage = StageValue
				continue
			} else if char == '-' {
				sign = -1
				stage = StageValue
				continue
			} else {
				stage = StageValue
			}
		}
		if stage == StageValue {
			if char > '9' || char < '0' {
				break
			} else {
				resultInt = resultInt*10 + int64(char-'0')
				if resultInt > math.MaxInt32{
					break
				}
			}
		}
	}
	tempResult := int64(sign) * resultInt
	if tempResult > math.MaxInt32 {
		return math.MaxInt32
	} else if tempResult < math.MinInt32 {
		return math.MinInt32
	} else {
		return int(tempResult)
	}
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 4321"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("9223372036854775808"))
}
