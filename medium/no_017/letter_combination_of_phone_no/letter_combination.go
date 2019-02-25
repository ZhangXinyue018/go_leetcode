package letter_combination_of_phone_no

var (
	charMap = map[int32][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
)

func letterCombinations(digits string) []string {
	result := []string{}
	for indexDigit, digit := range digits {
		temp := charMap[digit]
		if indexDigit == 0 {
			for _, value := range temp {
				result = append(result, string(value))
			}
		} else {
			lengthResult := len(result)
			for i := 0; i < lengthResult; i++ {
				for j := len(temp) - 1; j >= 0; j-- {
					if j == 0 {
						result[i] = result[i] + string(temp[j])
					} else {
						result = append(result, result[i]+string(temp[j]))
					}
				}
			}
		}
	}
	return result
}
