package main

import "fmt"

func searchInsert(nums []int, target int) int {
	return search(nums, target, 0, len(nums)-1)
}

func search(nums []int, target int, fromIndex int, toIndex int) int {
	if fromIndex == toIndex {
		if target > nums[fromIndex] {
			return fromIndex + 1
		} else {
			return fromIndex
		}
	} else {
		if target > nums[(fromIndex+toIndex)/2] {
			return search(nums, target, (fromIndex+toIndex)/2+1, toIndex)
		} else {
			return search(nums, target, fromIndex, (fromIndex+toIndex)/2)
		}
	}
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
	fmt.Println(searchInsert([]int{1,3,5,6}, 0))
}
