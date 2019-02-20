package balanced_binary_tree

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{&TreeNode{1,
			&TreeNode{2,
				&TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}},
				&TreeNode{3, nil, nil},
			},
			&TreeNode{2, nil, nil}},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
