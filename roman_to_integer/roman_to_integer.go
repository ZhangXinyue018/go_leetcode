package main

import "fmt"

var SymbolMap = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	length := len(s)
	result := SymbolMap[s[length-1]]
	for i := length - 2; i >= 0; i-- {
		if SymbolMap[s[i]] < SymbolMap[s[i+1]] {
			result = result - SymbolMap[s[i]]
		} else {
			result = result + SymbolMap[s[i]]
		}
	}
	return result
}

func main() () {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
