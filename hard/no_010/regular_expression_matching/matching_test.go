package regular_expression_matching

import "testing"

func TestMatching(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 1", args{"aa", "a"}, false},
		{"test 1", args{"aa", "a*"}, true},
		{"test 1", args{"ab", ".*"}, true},
		{"test 1", args{"aab", "c*a*b"}, true},
		{"test 1", args{"mississippi", "mis*is*p*."}, false},
		{"test 1", args{"aaa", "a*a"}, true},
		{"test 1", args{"aaa", "ab*a"}, false},
		{"test 1", args{"a", ".*..a*"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ITestMatching2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 1", args{"aa", "a"}, false},
		{"test 2", args{"aa", "a*"}, true},
		{"test 3", args{"ab", ".*"}, true},
		{"test 4", args{"aab", "c*a*b"}, true},
		{"test 5", args{"mississippi", "mis*is*p*."}, false},
		{"test 6", args{"aaa", "a*a"}, true},
		{"test 7", args{"aaa", "ab*a"}, false},
		{"test 8", args{"a", ".*..a*"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch2(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch2(%s, %s) = %v, want %v", tt.args.s, tt.args.p, got, tt.want)
			}
		})
	}
}
