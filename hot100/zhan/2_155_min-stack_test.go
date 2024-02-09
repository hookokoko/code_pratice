package zhan

import "testing"

type MinStack struct {
	data    []int
	minData []int // 栈内最小值, 长度跟data长度一致
}

func Constructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	var minData int
	if (len(this.minData) > 0 && val < this.minData[len(this.minData)-1]) || len(this.minData) == 0 {
		minData = val
	} else {
		minData = this.minData[len(this.minData)-1]
	}
	this.minData = append(this.minData, minData)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minData = this.minData[:len(this.minData)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[len(this.data)-1]
}

func TestGetMin(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin() //--> 返回 -3.
	minStack.Pop()
	minStack.Top()    // --> 返回 0.
	minStack.GetMin() // --> 返回 -2.
}
