package main

import (
	"testing"
)

func Test_strStr2(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{"a", "a"}, 0},
		{"test 2", args{"abaahaystacka", ""}, 0},
		{"test 3", args{"abaaa", "nb"}, -1},
		{"test 4", args{"jlalalalal", "al"}, 2},
		{"test 5", args{"a", "bababa"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr2(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{"a", "a"}, 0},
		{"test 2", args{"abaahaystacka", ""}, 0},
		{"test 3", args{"abaaa", "nb"}, -1},
		{"test 4", args{"jlalalalal", "al"}, 2},
		{"test 5", args{"a", "bababa"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
