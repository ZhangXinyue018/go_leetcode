package flatten_binary_tree

import (
	"fmt"
	"testing"
)

var input = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	},
	Right: nil,
}

func TestFlatten(t *testing.T) {
	fmt.Println("Start test 1")
	Flatten(input)
	temp := input
	for temp != nil {
		fmt.Println(temp.Val)
		fmt.Println(temp.Left)
		temp = temp.Right
	}
}

func TestFlatten2(t *testing.T) {
	fmt.Println("Start test 2")
	Flatten2(input)
	temp := input
	for temp != nil {
		fmt.Println(temp.Val)
		fmt.Println(temp.Left)
		temp = temp.Right
	}
}
