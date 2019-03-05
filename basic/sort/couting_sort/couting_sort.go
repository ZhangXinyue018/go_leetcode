package couting_sort

func CountingSort(nums []int) []int {
	max := 100
	countingMap := make([]int, max)
	for i := 0; i < len(nums); i++ {
		countingMap[nums[i]] = countingMap[nums[i]] + 1
	}
	result := []int{}
	for i := 0; i < max; i++ {
		for j := countingMap[i]; j > 0; j-- {
			result = append(result, i)
		}
	}
	return result
}
