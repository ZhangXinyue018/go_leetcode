package longest_valid_parethese

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	max := 0
	for i := 0; i <= len(s)-2; i++ {
		if s[i] == ')' && s[i+1] == '(' {
			left := maxWithFirst(s, i+1)
			right := maxWithEnd(s, i)
			sum := left + right
			if sum > max {
				max = sum
			}
		}
	}
	if s[0] == '(' {
		sum := maxWithFirst(s, 0)
		if sum > max {
			max = sum
		}
	}
	if s[len(s)-1] == ')' {
		sum := maxWithEnd(s, len(s)-1)
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxWithFirst(s string, j int) int {
	left := 0
	pointer := j - 1
	for i := j; i < len(s); i++ {
		char := s[i]
		if char == '(' {
			left++
		} else if left > 0 {
			left--
			if left == 0 {
				pointer = i
			}
		} else {
			break
		}
	}
	return pointer - j + 1
}

func maxWithEnd(s string, j int) int {
	right := 0
	pointer := j + 1
	for i := j; i >= 0; i-- {
		char := s[i]
		if char == ')' {
			right++
		} else if right > 0 {
			right--
			if right == 0 {
				pointer = i
			}
		} else {
			break
		}
	}
	return j - pointer + 1
}

func longestValidParenthesesBest(s string) int {
	if s == ""{
		return 0
	}
	max := 0
	leftList := []int{}
	encloseLength := make([]int, len(s))
	for i := 0; i< len(s);i++{
		char := s[i]
		if char == '('{
			leftList = append(leftList, i)
		}else if len(leftList) != 0{
			closeIndex := leftList[len(leftList)-1]
			leftList = leftList[:len(leftList)-1]
			currLength := i - closeIndex + 1
			if closeIndex > 0 && s[closeIndex-1] == ')'{
				currLength = currLength + encloseLength[closeIndex-1]
			}
			encloseLength[i] = currLength
			if currLength > max{
				max = currLength
			}
		}
	}
	return max
}
