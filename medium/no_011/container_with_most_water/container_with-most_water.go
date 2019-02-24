package container_with_most_water

func maxArea(height []int) int {
	left, right := 0 ,len(height)-1
	var max, square int
	for left < right {
		if height[left] > height [right] {
			square = height[right] * (right-left)
			right--
		} else {
			square = height[left] * (right-left)
			left++
		}
		if square > max {
			max = square
		}
	}
	return max
}
