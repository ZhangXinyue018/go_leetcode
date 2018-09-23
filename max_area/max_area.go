package main

import "fmt"

func maxArea(height []int) int {
	maxArea := 0
	heightLen := len(height)
	for index, heightInt := range height {
		for i := index + 1; i < heightLen; i++ {
			var area int
			if height[i] > heightInt {
				area = heightInt * (i - index)
			} else {
				area = height[i] * (i - index)
			}
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func maxAreaBetter(height []int) int {
	front := 0
	rear := len(height) - 1
	maxArea := 0
	for front < rear {
		var area int
		if height[front] < height[rear] {
			area = height[front] * (rear - front)
			front ++
		} else {
			area = height[rear] * (rear - front)
			rear--
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxAreaBetter([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxAreaBetter([]int{1, 2}))
}
