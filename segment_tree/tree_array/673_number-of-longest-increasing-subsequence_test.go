package tree_array

import (
	"fmt"
	"slices"
	"testing"
)

/*
给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
注意 这个数列必须是 严格 递增的。
*/

type Node struct {
	Max    int // 最长递增序列的长度
	MaxNum int // 最长递增序列的个数
}

type NumTree struct {
	Nums []Node
}

func NewNumTree(n int) NumTree {
	nodes := make([]Node, n+1)
	return NumTree{
		Nums: nodes,
	}
}

func (n *NumTree) query(idx int) Node {
	var ret Node
	for ; idx > 0; idx -= lowBit(idx) {
		ret = getMax(ret, n.Nums[idx])
	}
	return ret
}

func (n *NumTree) update(idx int, node Node) {
	for ; idx < len(n.Nums); idx += lowBit(idx) {
		n.Nums[idx] = getMax(n.Nums[idx], node)
		fmt.Printf("n.Nums[%d]>>>>%+v\n", idx, n.Nums[idx])
	}
}

func getMax(node1, node2 Node) Node {
	if node1.Max == node2.Max {
		return Node{
			Max:    node1.Max,
			MaxNum: node1.MaxNum + node2.MaxNum,
		}
	} else if node1.Max < node2.Max {
		return node2
	} else {
		return node1
	}
}

func findNumberOfLIS(nums []int) int {
	arrUniq := make([]int, len(nums))
	copy(arrUniq, nums)
	// 排序+去重
	slices.Sort(arrUniq)
	slices.Compact(arrUniq)

	tree := NewNumTree(len(arrUniq))
	ret := Node{}
	for i := 0; i < len(nums); i++ {
		idx := slices.Index(arrUniq, nums[i]) + 1
		node := tree.query(idx - 1) // 查找的是[0, i-1]的这个区间内的最大子序列长度及达到最大长度时的序列个数
		newNode := Node{
			Max:    node.Max + 1,        // 排序的话使得每遍历一个元素，都能够保证最长递增子序列的长度+1
			MaxNum: max(node.MaxNum, 1), // 个数最小也是1个，比如它自身
		}
		fmt.Printf("--------[i=%d]------[idx=%d]------node=%+v------new=%+v\n", i, idx, node, newNode)
		tree.update(idx, newNode)
		ret = getMax(ret, newNode)
	}
	return ret.MaxNum
}

func Test_FindNumberOfLIS(t *testing.T) {
	ret := findNumberOfLIS([]int{1, 3, 5, 4, 7})
	fmt.Println(ret)
}

//func Test_o(t *testing.T) {
//	fmt.Println(lowBit(5))
//}
