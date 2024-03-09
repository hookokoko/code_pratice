package shuangzhizhen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	s := newStack()
	s.Push(0)
	sum := 0
	for i, h := range height {
		if i == 0 {
			continue
		}
		if s.IsEmpty() || h <= height[s.Top()] {
			s.Push(i)
		} else {
			for !s.IsEmpty() && h >= height[s.Top()] {
				mid := s.Pop()
				if !s.IsEmpty() {
					left := s.Top()
					curSum := (min(h, height[left]) - height[mid]) * (i - left - 1)
					//fmt.Printf("curSum[%d]-----i[%d]----mid[%d]-----left[%d]\n", curSum, i, mid, left)
					sum = sum + curSum
				}
			}
			s.Push(i)
		}
	}
	return sum
}

func Test_trap(t *testing.T) {
	//assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))

}

type stack struct {
	data []int
}

func newStack() stack {
	return stack{data: make([]int, 0, 16)}
}
func (s *stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return res
}
func (s *stack) Push(val int) {
	s.data = append(s.data, val)
}
func (s *stack) Top() int {
	if s.IsEmpty() {
		return -1
	}
	return s.data[len(s.data)-1]
}
func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}
