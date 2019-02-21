package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	pointerA := 0
	pointerB := 0
	for pointerA < m || pointerB < n {
		if pointerA >= m+pointerB {
			nums1[m+pointerB] = nums2[pointerB]
			pointerB++
		} else if pointerB >= n {
			break
		} else {
			if nums1[pointerA] <= nums2[pointerB] {
				pointerA ++
			} else {
				for j := m + pointerB; j > pointerA; j-- {
					nums1[j] = nums1[j-1]
				}
				nums1[pointerA] = nums2[pointerB]
				pointerA++
				pointerB++
			}
		}
	}
	return
}

func mergeBetter(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	if n > 0 {
		copy(nums1[0:n], nums2)
	}
}
