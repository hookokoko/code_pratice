package binary_tree

import "code_pratice/common"

func RecursiveTraversal(root *common.Node[int]) []int {
	return nil
}

// NonRecursiveTraversal_Pre 先序
func NonRecursiveTraversal_Pre(root *common.Node[int]) []int {
	res := make([]int, 0, 32)
	stack := common.NewStack[*common.Node[int]]()

	if root == nil {
		return res
	}

	stack.Push(root)

	for !stack.IsEmpty() {
		node := stack.Top()
		if node != nil {
			res = append(res, node.Data)
		}
		stack.Pop()
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}

	return res
}

// NonRecursiveTraversal_Post 后序
func NonRecursiveTraversal_Post(root *common.Node[int]) []int {
	return nil
}

// NonRecursiveTraversal_In 中序
func NonRecursiveTraversal_In(root *common.Node[int]) []int {
	res := make([]int, 0, 32)
	stack := common.NewStack[*common.Node[int]]()
	cur := root // 标记左节点是否到头
	for cur != nil || !stack.IsEmpty() {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			cur = stack.Top()
			stack.Pop()
			res = append(res, cur.Data)
			cur = cur.Right
		}
	}
	return res
}
