package longest_palindrome

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{"test 1", args{""}, ""},
		//{"test 2", args{"babad"}, "bab"},
		//{"test 3", args{"cbbd"}, "bb"},
		{"test 4", args{"bananas"}, "anana"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindromeBetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{""}, ""},
		{"test 2", args{"babad"}, "bab"},
		{"test 3", args{"cbbd"}, "bb"},
		{"test 4", args{"a"}, "a"},
		{"test 5", args{"ccc"}, "ccc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeBetter(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
