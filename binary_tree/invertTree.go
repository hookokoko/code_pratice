package binary_tree

import "code_pratice/common"

func InvertTree(root *common.Node[int]) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
}
