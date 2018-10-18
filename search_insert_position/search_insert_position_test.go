package search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1,3,5,6}, 5}, 2},
		{"test 1", args{[]int{1,3,5,6}, 2}, 1},
		{"test 1", args{[]int{1,3,5,6}, 7}, 4},
		{"test 1", args{[]int{1,3,5,6}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}