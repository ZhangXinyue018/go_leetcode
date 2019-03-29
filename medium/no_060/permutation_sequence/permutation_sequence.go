package permutation_sequence

import "strconv"

func getPermutation(n int, k int) string {
	nums := []int{}
	indicator := 1
	for i := 1; i<=n;i++{
		nums = append(nums, i)
		indicator = indicator * i
	}
	return DetailPermutation(nums, k, "", indicator)
}

func DetailPermutation(nums []int, k int, pre string, indicator int)string{
	if len(nums) == 1 || k == 0{
		for i:=0;i<len(nums);i++{
			pre = pre + strconv.Itoa(nums[i])
		}
		return pre
	}
	indicator = indicator / len(nums)
	temp := (k-1) / indicator
	pre = pre +strconv.Itoa(nums[temp])
	k = k- temp * indicator
	newNums := append([]int{}, nums[:temp]...)
	newNums = append(newNums, nums[temp+1:]...)
	return DetailPermutation(newNums, k, pre, indicator)
}
