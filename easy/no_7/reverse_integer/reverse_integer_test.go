package reverse_integer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"test 1", 123, 321},
		{"test 2", 120, 21},
		{"test 3", -123, -321},
		{"test 4", 1534236469, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.x); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse2(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"test 1", 123, 321},
		{"test 2", 120, 21},
		{"test 3", -123, -321},
		{"test 4", 1534236469, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse2(tt.x); got != tt.want {
				t.Errorf("Reverse2() = %v, want %v", got, tt.want)
			}
		})
	}
}
