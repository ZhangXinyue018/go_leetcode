package merge_sort

import (
	"go_leetcode/basic/commons"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{4, 2, 3, 1}}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.nums); !commons.CheckArrayEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortLowMemVersion(t *testing.T) {
	type args struct {
		nums       []int
		leftIndex  int
		rightIndex int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]int{4, 2, 3, 1}, 0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSortLowMemVersion(tt.args.nums, tt.args.leftIndex, tt.args.rightIndex)
		})
	}
}
