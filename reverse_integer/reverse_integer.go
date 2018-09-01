package main

import (
		"math"
	"fmt"
)

func reverse(x int) int {
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

func reverse2(x int) int {
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

func main() () {
	fmt.Println(reverse2(123))
	fmt.Println(reverse2(120))
	fmt.Println(reverse2(-123))
	fmt.Println(reverse2(1534236469))
}
