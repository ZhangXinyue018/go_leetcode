package generate_parentheses

func generateParenthesis(n int) []string {
	root := &TestTreeNode{
		LeftCount:  n,
		RightCount: n,
		Value:      "",
		LeftNode:   nil,
		RightNode:  nil,
	}
	result := AppendString(root)
	return result

}

type TestTreeNode struct {
	LeftCount  int
	RightCount int
	Value      string
	LeftNode   *TestTreeNode
	RightNode  *TestTreeNode
}

func AppendString(root *TestTreeNode) []string {
	if root.LeftCount == 0 && root.RightCount == 0 {
		return []string{root.Value}
	}

	result := []string{}
	if root.RightCount > root.LeftCount {
		root.RightNode = &TestTreeNode{
			LeftCount:  root.LeftCount,
			RightCount: root.RightCount - 1,
			Value:      root.Value + ")",
		}
		result = append(result, AppendString(root.RightNode)...)
	}
	if root.LeftCount > 0 {
		root.LeftNode = &TestTreeNode{
			LeftCount:  root.LeftCount - 1,
			RightCount: root.RightCount,
			Value:      root.Value + "(",
		}
		result = append(result, AppendString(root.LeftNode)...)
	}
	return result
}
