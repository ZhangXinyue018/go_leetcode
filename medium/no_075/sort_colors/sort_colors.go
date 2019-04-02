package sort_colors

func sortColors(nums []int) {
	rpf := -1
	rpr := len(nums) - 1
	p := 0
	for p <= rpr {
		if nums[p] == 1 {
			if rpf == -1 {
				rpf = p
			}
			p++
		} else if nums[p] == 2 {
			nums[p], nums[rpr] = nums[rpr], nums[p]
			rpr--
		} else {
			if rpf == -1 {
				p++
			} else {
				nums[p], nums[rpf] = nums[rpf], nums[p]
				p++
				rpf++
			}
		}
	}
}
