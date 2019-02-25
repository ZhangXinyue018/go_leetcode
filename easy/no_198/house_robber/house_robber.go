package house_robber

func rob(nums []int) int {
	tempMap := make(map[int]int)

	for j := len(nums) - 1; j >= 0; j-- {
		if j == len(nums)-1 {
			tempMap[j] = nums[j]
		} else if j == len(nums)-2 {
			if nums[j] > nums[j+1] {
				tempMap[j] = nums[j]
			} else {
				tempMap[j] = nums[j+1]
			}
		} else {
			x := nums[j] + tempMap[j+2]
			y := tempMap[j+1]
			if x > y {
				tempMap[j] = x
			} else {
				tempMap[j] = y
			}
		}

	}

	return tempMap[0]
}
