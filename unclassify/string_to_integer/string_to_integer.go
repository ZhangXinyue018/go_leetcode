package string_to_integer

import (
	"math"
)

const (
	StageSign  = 1
	StageValue = 2
)

func MyAtoi(str string) int {
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
