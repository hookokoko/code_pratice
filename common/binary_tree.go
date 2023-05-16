package common

// RealNumber 支持对应数字类型及衍生数字类型
type RealNumber interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~int | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Node[T RealNumber] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

func NewNodeTree[T RealNumber](value T) *Node[T] {
	n := new(Node[T])
	n.Data = value
	n.Left = nil
	n.Right = nil
	return n
}

// Insert 二叉排序数的插入
func (root *Node[T]) Insert(value T) *Node[T] {
	if root == nil {
		return NewNodeTree[T](value)
	}
	if value < root.Data {
		root.Left = root.Left.Insert(value)
	} else {
		root.Right = root.Right.Insert(value)
	}
	return root
}

func (root *Node[T]) InsertList(valueList []T) *Node[T] {
	if len(valueList) == 0 {
		return root
	}
	for _, value := range valueList {
		root.Insert(value)
	}
	return root
}
