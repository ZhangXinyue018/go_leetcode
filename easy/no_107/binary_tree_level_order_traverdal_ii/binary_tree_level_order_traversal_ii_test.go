package binary_tree_level_order_traverdal_ii

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{&TreeNode{3,
			&TreeNode{9, nil, nil},
			&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}},
			[][]int{[]int{15, 7}, []int{9, 20}, []int{3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom2() = %v, want %v", got, tt.want)
			}
		})
	}
}
