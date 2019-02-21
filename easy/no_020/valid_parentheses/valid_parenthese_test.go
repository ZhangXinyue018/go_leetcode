package valid_parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s string
		want bool
	}{
		{"test 1", "()", true},
		{"test 2", "()[]{}", true},
		{"test 3", "(]", false},
		{"test 4", "([)]", false},
		{"test 5", "{[]}", true},
		{"test 6", "]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}