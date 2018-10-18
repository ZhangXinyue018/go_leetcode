package panlindrome_number

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"test 1", -123, false},
		{"test 1", 123, false},
		{"test 1", 121, true},
		{"test 1", 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
