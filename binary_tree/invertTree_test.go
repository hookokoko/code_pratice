package binary_tree

import (
	"code_pratice/common"
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	treeNode1 := common.NewNodeTree[int](20)
	treeNode1.Insert(10)
	treeNode1.Insert(30)
	treeNode1.Insert(40)
	treeNode1.Insert(22)

	InvertTree(treeNode1)

	fmt.Println(treeNode1)
}
