package binary_tree

import (
	"code_pratice/common"
	"fmt"
	"log"
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

func Test_BinaryTreePaths(t *testing.T) {
	root := common.NewTree([]string{"1", "2", "3", "4", "5", "6", "7"})
	res := BinaryTreePaths(&root)
	log.Println(res)
}
