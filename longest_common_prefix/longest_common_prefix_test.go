package longest_common_prefix

import (
	"testing"
)

func Test_longestCommonPrefix2(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"test 2", args{[]string{"dog", "racecar", "car"}}, ""},
		{"test 3", args{[]string{}}, ""},
		{"test 4", args{[]string{"aca", "cba"}}, ""},
		{"test 5", args{[]string{"aa", "ab"}}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix2(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"test 2", args{[]string{"dog", "racecar", "car"}}, ""},
		{"test 3", args{[]string{}}, ""},
		{"test 4", args{[]string{"aca", "cba"}}, ""},
		{"test 5", args{[]string{"aa", "ab"}}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
