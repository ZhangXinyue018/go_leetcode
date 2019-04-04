package remove_duplicate_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	i, p:= 0, 1
	pre, count := nums[0], 1
	for ;p<len(nums);p++{
		if nums[p] == pre {
			if count < 2{
				count++
				nums[i+1] = nums[p]
				i++
			}
		}else {
			pre = nums[p]
			count = 1
			nums[i+1] = nums[p]
			i++
		}
	}
	return i+1
}