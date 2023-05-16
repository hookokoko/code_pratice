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
