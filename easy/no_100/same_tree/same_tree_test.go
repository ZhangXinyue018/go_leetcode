package same_tree

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
			&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}}, true},
		{"test1", args{&TreeNode{1, &TreeNode{2, nil, nil}, nil},
			&TreeNode{1, nil, &TreeNode{2, nil, nil}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
