package max_area

import (
	"testing"
)

func Test_maxAreaBetter(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"test 2", args{[]int{1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaBetter(tt.args.height); got != tt.want {
				t.Errorf("maxAreaBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"test 2", args{[]int{1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
