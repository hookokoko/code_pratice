package common

import (
	"fmt"
	"testing"
)

func Test_Insert(t *testing.T) {
	root := NewNodeTree[int](123)
	root.Insert(24)
	fmt.Printf("%+v\n", root)
}

func TestNewTree(t *testing.T) {
	r := NewTree([]int{1, 2, 3, 4, 5, 6})
	t.Log(r)
}
