package two_sum

func TwoSum(nums []int, target int) []int {
	directionMap := map[int]int{}
	for index, value := range nums{
		oppIndex, hasKey := directionMap[value]
		if hasKey{
			return []int{oppIndex, index}
		}else{
			directionMap[target-value] = index
		}
	}
	return []int{}
}
