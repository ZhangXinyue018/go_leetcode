package recover_binary_search_tree

import (
	"testing"
)

func Test_recoverTree(t *testing.T) {
	root := &TreeNode{
		3,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{2, nil, nil}, nil}}

	t.Run("general test", func(t *testing.T) {
		recoverTree(root)
	})
	PrintPostOrder(root)
}
