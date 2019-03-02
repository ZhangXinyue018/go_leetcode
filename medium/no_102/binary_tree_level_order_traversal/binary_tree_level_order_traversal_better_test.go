package binary_tree_level_order_traversal

import (
	"fmt"
	"testing"
)

func Test_levelOrderBetter(t *testing.T) {
	root := &TreeNode{
		3,
		&TreeNode{1, nil, nil},
		&TreeNode{4, &TreeNode{2, nil, nil}, nil}}

	t.Run("general test", func(t *testing.T) {
		result := levelOrderBetter(root)
		fmt.Println(result)
	})
}
