package maximum_product_subarray

func maxProduct(nums []int) int {
	maxproduct := nums[0]
	length := len(nums)
	result := []int{}
	for i := 0; i < length; i++ {
		result[i] = 1
	}
	for i := 1; i <= length; i++ {
		for j := 0; j <= length-i; j++ {
			result[j] = result[j] * nums[i-1+j]
			if result[j] > maxproduct {
				maxproduct = result[j]
			}
		}
	}
	return maxproduct
}

func maxProductImprove(nums []int) int {
	var max, min, result int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max = nums[0]
			min = nums[0]
			result = nums[0]
		} else {
			newMax := getMax(max*nums[i], min*nums[i], nums[i])
			newMin := getMin(min*nums[i], max*nums[i], nums[i])
			if newMax > result {
				result = newMax
			}
			max = newMax
			min = newMin
		}
	}
	return result
}

func getMax(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	} else {
		if b > c {
			return b
		}
		return c
	}
}

func getMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	} else {
		if b < c {
			return b
		}
		return c
	}
}
