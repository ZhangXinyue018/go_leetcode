package string_to_integer

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"test 1", "42", 42},
		{"test 1", "   -42", -42},
		{"test 1", "4193 with words", 4193},
		{"test 1", "words and 4321", 0},
		{"test 1", "-91283472332", math.MinInt32},
		{"test 1", "9223372036854775808", math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAtoi(tt.str); got != tt.want {
				t.Errorf("MyAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyAtoi2(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"test 1", "42", 42},
		{"test 1", "   -42", -42},
		{"test 1", "4193 with words", 4193},
		{"test 1", "words and 4321", 0},
		{"test 1", "-91283472332", math.MinInt32},
		{"test 1", "9223372036854775808", math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAtoi2(tt.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestMyAtoi3(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"test 1", "42", 42},
		{"test 1", "   -42", -42},
		{"test 1", "4193 with words", 4193},
		{"test 1", "words and 4321", 0},
		{"test 1", "-91283472332", math.MinInt32},
		{"test 1", "9223372036854775808", math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MyAtoi3(tt.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}