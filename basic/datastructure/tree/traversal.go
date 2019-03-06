package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InOrderTraversal(root *TreeNode) () {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	fmt.Println(root.Val)
	InOrderTraversal(root.Right)
}

func PreOrderTraversal(root *TreeNode) () {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PostOrderTraversal(root.Right)
	PostOrderTraversal(root.Left)
}
