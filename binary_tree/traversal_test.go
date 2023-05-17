package binary_tree

import (
	"code_pratice/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonRecursiveTraversal_Pre(t *testing.T) {
	treeNode := common.NewNodeTree[int](20)
	treeNode.Insert(10)
	treeNode.Insert(30)
	//treeNode.Insert(40)
	//treeNode.Insert(33)
	//treeNode.Insert(32)
	//treeNode.Insert(69)
	res := NonRecursiveTraversal_Pre(treeNode)
	assert.Equal(t, []int{20, 10, 30}, res)
}

func TestLevelOrder(t *testing.T) {
	treeNode1 := common.NewNodeTree[int](20)
	treeNode1.Insert(10)
	treeNode1.Insert(30)
	treeNode1.Insert(40)
	treeNode1.Insert(22)

	res := LevelOrder(treeNode1)
	assert.Equal(t, [][]int{{20}, {10, 30}, {22, 40}}, res)

	treeNode2 := common.NewNodeTree[int](20)

	res2 := LevelOrder(treeNode2)
	assert.Equal(t, [][]int{{20}}, res2)

	//nilInt := new(int)
	//treeNode3 := common.NewNodeTree[int](*nilInt)
	//
	//res3 := LevelOrder(treeNode3)
	//assert.Equal(t, [][]int{}, res3)
}
