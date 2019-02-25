package _sum

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{"test1", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{
		//	{2, 1, -1, -2}, {2, 0, 0, -2}, {1, 0, 0, -1}}},
		{"test2", args{[]int{0, 0, 0, 0}, 1}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
