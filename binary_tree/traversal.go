package binary_tree

import (
	"code_pratice/common"
	"container/list"
)

func RecursiveTraversal(root *common.Node[int]) []int {
	return nil
}

// NonRecursiveTraversal_Pre 先序
func NonRecursiveTraversal_Pre(root *common.Node[int]) []int {
	res := make([]int, 0, 32)
	stack := common.NewStack[*common.Node[int]](32)

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
	stack := common.NewStack[*common.Node[int]](32)
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

// LevelOrder 层次遍历
func LevelOrder(root *common.Node[int]) [][]int {
	if root == nil {
		return [][]int{}
	}
	l := list.New()
	res := make([][]int, 0, 32)

	// 入队
	l.PushBack(root)
	for l.Len() > 0 {
		size := l.Len()
		sub_res := make([]int, 0, 32)
		for i := 0; i < size; i++ {
			// 出队
			temp := l.Front()
			l.Remove(temp)

			left := temp.Value.(*common.Node[int]).Left
			right := temp.Value.(*common.Node[int]).Right
			if left != nil {
				l.PushBack(left)
			}
			if right != nil {
				l.PushBack(right)
			}

			sub_res = append(sub_res, temp.Value.(*common.Node[int]).Data)
		}
		res = append(res, sub_res)
	}
	return res
}
