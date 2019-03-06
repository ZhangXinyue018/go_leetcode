package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	t.Run("test insert", func(t *testing.T) {
		tree := &TreeNode{7, nil, nil}
		Insert(tree, 1)
		Insert(tree, 6)
		Insert(tree, 4)
		Insert(tree, 9)
		Insert(tree, 20)
		Insert(tree, 17)
		Insert(tree, 18)
		Insert(tree, 5)
		Insert(tree, 13)
		Insert(tree, 2)
		InOrderTraversal(tree)
	})

}

func TestDelete(t *testing.T) {
	t.Run("test delete", func(t *testing.T) {
		tree := &TreeNode{7, nil, nil}
		Insert(tree, 1)
		Insert(tree, 6)
		Insert(tree, 4)
		Insert(tree, 9)
		Insert(tree, 20)
		Insert(tree, 17)
		Insert(tree, 18)
		Insert(tree, 5)
		Insert(tree, 13)
		Insert(tree, 2)
		InOrderTraversal(tree)
		fmt.Println("=====================")
		Delete(tree, 20)
		InOrderTraversal(tree)
	})
}
