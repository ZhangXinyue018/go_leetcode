package cutting_trees

func Solution(nums []int) (int) {
	temp := []int{}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			temp = append(temp, i)
		}
	}
	if len(temp) == 0 {
		return len(nums)
	} else if len(temp) > 1 {
		return 0
	}

	index := temp[0]
	count := 2
	if len(nums) > index+2 && nums[index] > nums[index+2] {
		count --
	}
	if index != 0 && nums[index-1] > nums[index+1] {
		count--
	}
	return count
}
