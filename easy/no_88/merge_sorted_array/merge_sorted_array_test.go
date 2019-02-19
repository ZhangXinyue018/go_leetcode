package merge_sorted_array

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}},
		{"test2", args{[]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			fmt.Println(tt.args.nums1)
			//if [...]int(tt.args.nums1) != [...]int(tt.want) {
			//	t.Errorf("deleteDuplicates() = %v, want %v", tt.args.nums1, tt.want)
			//}
		})
	}
}

func Test_mergeBetter(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}},
		{"test2", args{[]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeBetter(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			fmt.Println(tt.args.nums1)
		})
	}
}
