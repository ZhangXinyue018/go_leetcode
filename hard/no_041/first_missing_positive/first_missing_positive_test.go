package first_missing_positive

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{nums: []int{1}}, 2},
		{"test2", args{nums: []int{7, 8, 9, 11, 12}}, 1},
		{"test3", args{nums: []int{3, 4, -1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("%s: firstMissingPositive() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
