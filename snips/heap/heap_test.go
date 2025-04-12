package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestSmallHeap(t *testing.T) {
	sh := SmallHeap{Value: make([][2]int, 0)}
	sh.Value = [][2]int{
		{1, 6},
		{3, 2},
		{5, 3},
		{7, 1},
	}
	heap.Init(&sh)
	// 依次取出堆顶元素
	for sh.Len() > 0 {
		item := heap.Pop(&sh).([2]int)
		fmt.Printf("%v ", item)
	}
}
