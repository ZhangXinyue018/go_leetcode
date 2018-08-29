package main

import "fmt"

// ************* Problem #4 *************

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	pointer1 := 0
	length1 := len(nums1)
	pointer2 := 0
	length2 := len(nums2)
	result := make([]int, 0)
	for pointer1 < length1 && pointer2 < length2 {
		if nums1[pointer1] <= nums2[pointer2] {
			result = append(result, nums1[pointer1])
			pointer1++
		} else {
			result = append(result, nums2[pointer2])
			pointer2++
		}
	}
	for pointer1 < length1 {
		result = append(result, nums1[pointer1])
		pointer1++
	}
	for pointer2 < length2 {
		result = append(result, nums2[pointer2])
		pointer2++
	}
	totalLength := length1 + length2
	if totalLength%2 == 0 {
		return float64(result[totalLength/2-1]+result[totalLength/2]) / float64(2)
	} else {
		return float64(result[totalLength/2])
	}
}

func main() () {
	num1 := []int{1, 2, 3}
	num2 := []int{5, 6}
	fmt.Println(FindMedianSortedArrays(num1, num2))
}
