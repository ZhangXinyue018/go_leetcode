package contruct_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	remarkList := make([]bool, len(preorder))
	return DetailBuildTree(preorder, &remarkList, inorder, 0, len(preorder)-1)

}

func DetailBuildTree(preorder []int, remarkList *[]bool, inorder []int, leftIndex, rightIndex int) *TreeNode {
	if leftIndex > rightIndex {
		return nil
	}
	i := 0
	for ; i < len(preorder); i++ {
		if !(*remarkList)[i] {
			break
		}
	}
	val := preorder[i]
	(*remarkList)[i] = true
	j := leftIndex
	for ; j <= rightIndex; j++ {
		if inorder[j] == val {
			break
		}
	}

	return &TreeNode{
		Val:   val,
		Left:  DetailBuildTree(preorder, remarkList, inorder, leftIndex, j-1),
		Right: DetailBuildTree(preorder, remarkList, inorder, j+1, rightIndex),
	}
}
