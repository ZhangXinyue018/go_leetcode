package recover_binary_search_tree

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var FirstHit, LastHit *TreeNode
	PreScan := &TreeNode{math.MinInt64, nil, nil}
	FirstHit, LastHit, PreScan = FillTheTwoHit(root, FirstHit, LastHit, PreScan)
	temp := FirstHit.Val
	FirstHit.Val = LastHit.Val
	LastHit.Val = temp
}

func FillTheTwoHit(root *TreeNode, FirstHit *TreeNode, LastHit *TreeNode,
	PreScan *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if root != nil {
		FirstHit, LastHit, PreScan = FillTheTwoHit(root.Left, FirstHit, LastHit, PreScan)
		if root.Val < PreScan.Val {
			if FirstHit == nil {
				FirstHit = PreScan
			}
			LastHit = root
		}
		PreScan = root
		FirstHit, LastHit, PreScan = FillTheTwoHit(root.Right, FirstHit, LastHit, PreScan)
	}
	return FirstHit, LastHit, PreScan
}

func PrintInOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	PrintInOrder(root.Left)
	fmt.Println(root.Val)
	PrintInOrder(root.Right)
}

func PrintPostOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(root.Val)
	PrintPostOrder(root.Left)
	PrintPostOrder(root.Right)
}
