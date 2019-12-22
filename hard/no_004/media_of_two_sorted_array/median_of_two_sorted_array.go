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


func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m < n {
		return findMedianSortedArrays2(nums2, nums1)
	}
	if n == 0{
		return float64(nums1[m/2] + nums1[(m-1)/2])/2.0
	}
	if m == 1 && n == 1{
		return float64(nums1[0] + nums2[0])/2.0
	}
	if n == 1 || n == 2 {
		nums1s := 0
		nums1e := m
		nums2s := 0
		nums2e := n
		if nums1[0] < nums2[0]{
			nums1s = 1
		}else{
			nums2s = 1
		}
		if nums1[m-1] > nums2[n-1]{
			nums1e = m-1
		}else{
			nums2e = n-1
		}
		return findMedianSortedArrays2(nums1[nums1s:nums1e], nums2[nums2s:nums2e])
	}

	if nums1[m/2] < nums2[n/2]{
		return findMedianSortedArrays2(nums1[(n-1)/2:], nums2[:n-(n-1)/2])
	}else{
		return findMedianSortedArrays2(nums1[:m-(n-1)/2], nums2[(n-1)/2:])
	}
}
