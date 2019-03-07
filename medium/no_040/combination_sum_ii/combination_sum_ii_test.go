package combination_sum_ii

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{"test1", args{[]int{10, 1, 2, 7, 6, 1, 5}, 8}, [][]int{
		//	{1, 2, 5},
		//	{1, 7},
		//	{1, 1, 6},
		//	{2, 6},
		//}},
		{"test2", args{[]int{2, 5, 2, 1, 2}, 5}, [][]int{
			{5}, {1, 2, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combinationSum22(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{10, 1, 2, 7, 6, 1, 5}, 8}, [][]int{
			{1, 2, 5},
			{1, 7},
			{1, 1, 6},
			{2, 6},
		}},
		{"test2", args{[]int{2, 5, 2, 1, 2}, 5}, [][]int{
			{5}, {1, 2, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum22(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum22() = %v, want %v", got, tt.want)
			}
		})
	}
}
