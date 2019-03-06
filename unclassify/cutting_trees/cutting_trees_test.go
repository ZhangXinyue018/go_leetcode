package cutting_trees

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 4, 5, 4}}, 2},
		{"test2", args{[]int{1, 2, 3, 4}}, 4},
		{"test3", args{[]int{2, 5, 4}}, 2},
		{"test4", args{[]int{5, 3, 4}}, 1},
		{"test5", args{[]int{5, 6, 5, 6, 5}}, 0},
		{"test6", args{[]int{4, 5, 2, 3, 4}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.nums); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
