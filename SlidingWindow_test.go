package code_pratice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 	1. 实现滑动窗口的泛型
	2. 具体自定义的push、pop方法，不仅仅局限于 数值比较
*/

type MyQueue[T any] struct {
	queue []T
	Comparable
}

type Comparable interface {
	CanPush(val, back any) bool
	CanPop(val, front any) bool
}

func NewMyQueue[T any](comp Comparable) *MyQueue[T] {
	return &MyQueue[T]{
		queue:      make([]T, 0),
		Comparable: comp,
	}
}

func (m *MyQueue[T]) Front() T {
	return m.queue[0]
}

func (m *MyQueue[T]) Back() T {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue[T]) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue[T]) Push(val T) {
	//for !m.Empty() && val > m.Back() {
	//	m.queue = m.queue[:len(m.queue)-1]
	//}
	for !m.Empty() && m.CanPush(val, m.Back()) {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue[T]) Pop(val int) {
	//if !m.Empty() && val == m.Front() {
	//	m.queue = m.queue[1:]
	//}
	if !m.Empty() && m.CanPop(val, m.Front()) {
		m.queue = m.queue[1:]
	}
}

type comp struct{}

func (c comp) CanPush(val, back any) bool {
	return val.(int) > back.(int)
}

func (c comp) CanPop(val, front any) bool {
	return val == front
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	q := NewMyQueue[int](comp{})
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	// 需要记录下前k个元素最大值
	res = append(res, q.Front())
	for i := k; i < len(nums); i++ {
		// 注意单调队列 入队、出队顺序
		q.Pop(nums[i-k])
		q.Push(nums[i])
		res = append(res, q.Front())
	}
	return res
}

func Test_Result(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := maxSlidingWindow(tt.nums, tt.k)
			assert.Equal(t, tt.want, res)
		})
	}
}
