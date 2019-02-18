package max_subarray

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{-2, 1}, 1},
		{"test 2", []int{2, 3, 4}, 9},
		{"test 3", []int{-2, 3, 4}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test 1", []int{-2, 1}, 1},
		{"test 2", []int{2, 3, 4}, 9},
		{"test 3", []int{-2, 3, 4}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray2(tt.nums); got != tt.want {
				t.Errorf("maxSubArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}
