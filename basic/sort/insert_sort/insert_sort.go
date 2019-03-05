package insert_sort

// 插入排序可以归纳为增量法

func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
	return nums
}

func InsertSortAnother(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}
