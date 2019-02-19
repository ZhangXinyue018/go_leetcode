package zigzag_conversion

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"test 2", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"test 3", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"test 4", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"test 5", args{"AB", 1}, "AB"},
		{"test 6", args{"AB", 1}, "AB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertBetter(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"test 2", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"test 3", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"test 4", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"test 5", args{"AB", 1}, "AB"},
		{"test 6", args{"AB", 1}, "AB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertBetter(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("ConvertBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

