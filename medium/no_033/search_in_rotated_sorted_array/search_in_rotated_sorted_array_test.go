package search_in_rotated_sorted_array

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"test2", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchBetter(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"test2", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{"test3", args{[]int{5, 1, 3}, 3}, 2},
		{"test4", args{[]int{1, 3}, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBetter(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
