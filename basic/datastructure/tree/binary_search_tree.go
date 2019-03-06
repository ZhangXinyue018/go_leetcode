package tree

func Insert(root *TreeNode, num int) *TreeNode {
	temp := root
	for temp != nil {
		if num > temp.Val {
			if temp.Right == nil {
				temp.Right = &TreeNode{num, nil, nil}
				return root
			}
			temp = temp.Right
		} else {
			if temp.Left == nil {
				temp.Left = &TreeNode{num, nil, nil}
				return root
			}
			temp = temp.Left
		}
	}
	return &TreeNode{num, nil, nil}
}

func SearchNode(root *TreeNode, num int) (*TreeNode) {
	if root == nil {
		return nil
	} else if root.Val == num {
		return root
	}
	temp := root
	for temp != nil {
		if num > temp.Val {
			temp = temp.Right
		} else if num < temp.Val {
			temp = temp.Left
		} else {
			return temp
		}
	}
	return temp
}

func Delete(root *TreeNode, num int) *TreeNode {
	tempNode := SearchNode(root, num)
	if tempNode.Left == nil {
		*tempNode = *tempNode.Right
	} else if tempNode.Right == nil {
		*tempNode = *tempNode.Left
	} else {
		minOnRight := FindMin(tempNode.Right)
		tempNode.Val = minOnRight.Val
		*minOnRight = *minOnRight.Right
	}
	return root
}

func FindMin(root *TreeNode) (*TreeNode) {
	temp := root
	for temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func FindMax(root *TreeNode) (*TreeNode) {
	temp := root
	for temp.Right != nil {
		temp = temp.Right
	}
	return temp
}
