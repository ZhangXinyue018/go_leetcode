package majority_element

func majorityElement(nums []int) int {
	resultMap := make(map[int]int)
	length := len(nums)
	for _, value := range nums {
		if _, ok := resultMap[value]; ok {
			resultMap[value] = resultMap[value] + 1
		} else {
			resultMap[value] = 1
		}
		if resultMap[value] > length/2 {
			return value
		}
	}
	return 0
}

func majorityElementBetter(nums []int) int {
	count := 1
	num := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == num {
			count++
		} else {
			count--
		}
		if count == 0 {
			num = nums[i+1]
		}
	}
	return num
}
