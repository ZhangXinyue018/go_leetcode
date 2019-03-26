package jump_game

func canJump(nums []int) bool {
	if len(nums) == 1{
		return true
	}
	target := 1
	i := len(nums)-2
	for ; i >=0;i--{
		if nums[i]>=target{
			break
		}
		target++
	}
	if i <0{
		return false
	}else{
		return canJump(nums[:i+1])
	}
}
