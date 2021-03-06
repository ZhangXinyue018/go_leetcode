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
				if resultInt > math.MaxInt32 {
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

func MyAtoi2(str string) int {
	sign := 0
	var result int64 = 0

	for _, char := range str {
		if sign == 0 {
			if char == '+' {
				sign = 1
			} else if char == '-' {
				sign = -1
			} else if char != ' ' {
				temp := char - '0'
				if temp < 0 || temp > 9 {
					return 0
				}
				sign = 1
				result = int64(temp)
			}
		} else {
			temp := char - '0'
			if temp < 0 || temp > 9 {
				break
			}
			result = result*int64(10) + int64(temp)
			if result > math.MaxInt32{
				break
			}
		}
	}

	result = result * int64(sign)
	if result > math.MaxInt32{
		return math.MaxInt32
	}else if result < math.MinInt32{
		return math.MinInt32
	}else{
		return int(result)
	}

}

func MyAtoi3(str string) int {
	if len(str) >0 && str[0] == ' '{
		return MyAtoi3(str[1:])
	}else if len(str) == 0{
		return 0
	}
	sig := int64(1)
	var result int64
	for i :=0 ;i < len(str); i++{
		if i == 0 {
			if str[i] == '+'{
				continue
			}else if str[i] == '-'{
				sig = int64(-1)
				continue
			}
		}
		if str[i] >= '0' && str[i] <= '9' && result <= math.MaxInt32{
			result = 10 * result + int64(str[i] - '0')
		}else{
			break
		}
	}
	result = sig * result
	if result < math.MinInt32{
		return math.MinInt32
	}else if result > math.MaxInt32{
		return math.MaxInt32
	}else{
		return int(result)
	}
}

