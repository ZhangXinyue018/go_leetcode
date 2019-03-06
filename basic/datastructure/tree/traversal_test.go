package tree

import "testing"

var root = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{3, nil, nil},
		Right: &TreeNode{4, nil, nil},
	},
	Right: &TreeNode{
		Val:   5,
		Left:  &TreeNode{6, nil, nil},
		Right: &TreeNode{7, nil, nil},
	},
}

func TestInOrderTraversal(t *testing.T) {
	t.Run("test in order", func(t *testing.T) {
		InOrderTraversal(root)
	})
}

func TestPreOrderTraversal(t *testing.T) {
	t.Run("test pre order", func(t *testing.T) {
		PreOrderTraversal(root)
	})
}

func TestPostOrderTraversal(t *testing.T) {
	t.Run("test post order", func(t *testing.T) {
		PostOrderTraversal(root)
	})
}
