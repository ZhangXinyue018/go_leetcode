package int_to_roman

import (
	"testing"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{8}, "VIII"},
		{"test 2", args{3}, "III"},
		{"test 3", args{4}, "IV"},
		{"test 4", args{9}, "IX"},
		{"test 5", args{58}, "LVIII"},
		{"test 6", args{1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	testInt := 1994
	resultStr := "MCMXCIV"

	for i := 0; i < b.N; i++ {
		actual := intToRoman(testInt)
		if actual != resultStr {
			b.Errorf("intToRoman() = %v, want %v", actual, resultStr)
		}
	}
}

func Test_intToRomanTest(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 1", args{8}, "VIII"},
		{"test 2", args{3}, "III"},
		{"test 3", args{4}, "IV"},
		{"test 4", args{9}, "IX"},
		{"test 5", args{58}, "LVIII"},
		{"test 6", args{1994}, "MCMXCIV"},
		{"test 7", args{20}, "XX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRomanTest(tt.args.num); got != tt.want {
				t.Errorf("intToRomanTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
