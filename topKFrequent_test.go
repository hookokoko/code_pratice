package code_pratice

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
 1. 统计数组中对应数字的频率到map
 2. 对这些数组根据频率存入堆中，去k个堆顶的值行
 这题主要想联系下golang中heap的实现
*/

type Item struct {
	value     int // 值
	frequency int // 出现频率优先级
}

type Items []Item

// 这个就是将元素加入到队列，head会帮我调整堆顶
func (items *Items) Push(x any) {
	*items = append(*items, x.(Item))
}

// 这个就是将元素从队列中去掉，head会帮我调整堆顶
func (items *Items) Pop() any {
	var v Item
	*items, v = (*items)[:items.Len()-1], (*items)[items.Len()-1]
	return v
}

func (items *Items) Len() int {
	return len(*items)
}

func (items *Items) Less(i, j int) bool {
	// 注意这里i、j的顺序，顺序不同对应的就是大顶堆和小顶堆的区别
	return (*items)[i].frequency > (*items)[j].frequency
}

func (items *Items) Swap(i, j int) {
	(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
}

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	items := new(Items)
	for num, freq := range m {
		*items = append(*items, Item{
			value:     num,
			frequency: freq,
		})
	}

	heap.Init(items)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(items).(Item).value)
	}

	return res
}

func Test_TopKFrequent(t *testing.T) {
	res := topKFrequent([]int{1, 3, 2, 2, 1, 1}, 2)
	assert.Equal(t, []int{1, 2}, res)
}
