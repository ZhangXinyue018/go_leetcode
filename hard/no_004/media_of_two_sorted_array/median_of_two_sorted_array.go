package media_of_two_sorted_array

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 0 {
		a := SearchK(nums1, nums2, totalLen/2)
		b := SearchK(nums1, nums2, totalLen/2-1)
		return (float64(a) + float64(b)) / 2
	} else {
		a := SearchK(nums1, nums2, totalLen/2)
		return float64(a)
	}
}

func SearchK(m []int, n []int, k int) int {
	if len(m) == 0 {
		return n[k]
	} else if len(n) == 0 {
		return m[k]
	}

	halfNIndex := len(n) / 2
	if m[0] >= n[halfNIndex] {
		if k == halfNIndex {
			return n[halfNIndex]
		} else if k > halfNIndex {
			return SearchK(m, n[halfNIndex+1:], k-halfNIndex-1)
		} else {
			return SearchK([]int{}, n[:halfNIndex], k)
		}
	} else {
		if m[0] < n[0] {
			if k == 0 {
				return m[0]
			}
			return SearchK(m[1:], n, k-1)
		} else {
			if k == 0 {
				return n[0]
			}
			return SearchK(m, n[1:], k-1)
		}
	}
}
