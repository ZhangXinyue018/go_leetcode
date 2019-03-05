package couting_sort

import (
	"go_leetcode/basic/commons"
	"testing"
)

func TestCountingSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{4, 2, 3, 1}}, []int{1, 2, 3, 4}},
		{"test2", args{[]int{3, 1, 4, 2}}, []int{1, 2, 3, 4}},
		{"test3", args{[]int{2, 8, 7, 1, 3, 5, 6, 4}}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountingSort(tt.args.nums); !commons.CheckArrayEqual(got, tt.want) {
				t.Errorf("CoutingSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
