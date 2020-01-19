package trapping_rain_water

func trap(height []int) int {
	return trapDetail(height, 0, len(height)-1)
}

func trapDetail(height []int, from, to int) int {
	if from >= to-1 {
		return 0
	}
	areaAcc := 0
	heightLine, maxLine := minAndMax(height[from], height[to])
	maxValue := 0
	maxIndex := -1
	for i := from + 1; i < to; i++ {
		if height[i] > maxLine {
			return trapDetail(height, from, i) + trapDetail(height, i, to)
		}
		if height[i] > maxValue {
			maxValue = height[i]
			maxIndex = i
		}
		areaAcc = areaAcc + height[i]
	}
	if maxValue > heightLine {
		return trapDetail(height, from, maxIndex) + trapDetail(height, maxIndex, to)
	}
	return (to-from-1)*heightLine - areaAcc
}

func minAndMax(x, y int) (int, int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}
