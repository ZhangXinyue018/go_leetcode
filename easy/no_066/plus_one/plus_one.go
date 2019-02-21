package plus_one

func plusOne(digits []int) []int {
	cin := 1
	for j := len(digits) - 1; j >= 0; j-- {
		temp := digits[j] + cin
		if temp < 10 {
			digits[j] = temp
			cin = 0
			break
		} else {
			digits[j] = temp % 10
			cin = 1
		}
	}
	if cin == 0{
		return digits
	}else{
		return append([]int{1}, digits...)
	}
}
