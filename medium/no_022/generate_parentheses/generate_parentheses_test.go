package generate_parentheses

import (
	"fmt"
	"reflect"
	"testing"
)

func ITest_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{1}, []string{"()"}},
		{"test2", args{2}, []string{"()()", "(())"}},
		{"test3", args{3}, []string{"()()()", "()(())", "(())()", "(()())", "((()))"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateParenthesisBest(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test1", args{1}, []string{"()"}},
		{"test2", args{2}, []string{"()()", "(())"}},
		{"test3", args{3}, []string{"()()()", "()(())", "(())()", "(()())", "((()))"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(generateParenthesisBest(tt.args.n))
			//if got := generateParenthesisBest(tt.args.n); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			//}
		})
	}
}
