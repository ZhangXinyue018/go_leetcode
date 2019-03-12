package longest_common_string

import "testing"

func TestLongestCommonSubSequence(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abc", "cabc"}, 3},
		{"test2", args{"abc", "atc"}, 2}, //"ac" is consider as common in sub sequence
		{"test3", args{"abc", "gvt"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonSubSequence(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LongestCommonSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
