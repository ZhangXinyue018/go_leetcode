package longest_substring_without_repeating_characters

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{"a"}, 1},
		{"test 2", args{"bbbb"}, 1},
		{"test 3", args{"pwwkew"}, 3},
		{"test 4", args{"ababab"}, 2},
		{"test 5", args{"abba"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
