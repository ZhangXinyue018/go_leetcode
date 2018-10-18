package divide_two_integers

import (
		"math"
)

func divide(dividend int, divisor int) int {
	if dividend > math.MaxInt32 || dividend < math.MinInt32 {
		return math.MaxInt32
	}
	sign := 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}
	result := 0
	for dividend >= divisor {
		temp, multiple := divisor, 1
		for dividend >= (temp << 1) {
			temp = temp << 1
			multiple = multiple << 1
		}
		dividend = dividend - temp
		result = result + multiple
	}
	if sign > 0 {
		if result > math.MaxInt32 {
			return math.MaxInt32
		}
		return result
	} else {
		if -result < math.MinInt32 {
			return math.MinInt32
		}
		return -result
	}
}
