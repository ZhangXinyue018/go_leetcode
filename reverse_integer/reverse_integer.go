package reverse_integer

import (
	"math"
)

func Reverse(x int) int {
	var input int64
	var result int64 = 0
	if x > 0 {
		input = int64(x)
		for input > 0 {
			result = result*10 + input%10
			input = input / 10
		}
	} else {
		input = -int64(x)
		for input > 0 {
			result = result*10 + input%10
			input = input / 10
		}
		result = -result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return int(result)
}

func Reverse2(x int) int {
	var result int64 = 0
	var input = int64(x)
	for input != 0 {
		result = result*10 + input%10
		input = input / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return int(result)
}
