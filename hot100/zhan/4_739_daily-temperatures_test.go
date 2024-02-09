package zhan

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type stack []int

func (s *stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *stack) Push(val int) {
	*s = append(*s, val)
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	//fmt.Println(res)
	slice := make([]int, len(temperatures))
	st := stack(slice)
	// i从1开始，因为比较i和i-1
	for i := 1; i < len(temperatures); i++ {
		// 如果小于等于，就入栈。相当于单调递减栈，这样就能找到第一个大于等于的当前i的元素距离
		if temperatures[i] <= temperatures[st.Top()] {
			st.Push(i)
		} else {
			// 如果碰见了大于（递减栈）栈顶的元素，就是一直弹出，直到栈中的元素都大于当前i的值；
			// 其中，每弹出一个值，即可保存这个值的距离了
			for !st.IsEmpty() && temperatures[i] > temperatures[st.Top()] {
				res[st.Top()] = i - st.Top()
				st.Pop()
			}
			// 需要把这个i再入栈，维持住单调递减的特性
			st.Push(i)
		}
	}
	return res
}

func TestDailyTemperatures(t *testing.T) {
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
