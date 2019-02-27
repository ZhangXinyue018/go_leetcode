package permutation

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{"abc"}, []string{
			"abc", "acb",
			"bac", "bca",
			"cab", "cba",
	}},
		{"test2", args{"abcd"}, []string{
			"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
			"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
			"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
			"dabc", "dacb", "dbac", "dbca", "dcab", "dcba",
	}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
