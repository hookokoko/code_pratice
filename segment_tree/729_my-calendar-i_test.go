package segment_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 常量N意味小数组的性能差
const N = 1e2

type INode struct {
	left  *INode
	right *INode
	val   int
	lazy  int
}

func (in *INode) query(start, end, left, right int) int {
	if start >= left && end <= right {
		return in.val
	}
	mid := start + (end-start)>>1
	ans := 0
	in.pushDown(mid-start+1, end-mid)
	if left <= mid {
		ans = in.left.query(start, mid, left, right)
	}
	if right > mid {
		ans = max(ans, in.right.query(mid+1, end, left, right))
	}
	return ans
}

func (in *INode) update(start, end, left, right, val int) {
	if start >= left && end <= right {
		in.val += val // 这里不用val*节点数
		in.lazy += val
		return
	}
	mid := start + (end-start)>>1
	in.pushDown(mid-start+1, end-mid)
	if left <= mid {
		in.left.update(start, mid, left, right, val)
	}
	if right > mid {
		in.right.update(mid+1, end, left, right, val)
	}
	in.pushUp()
}

func (in *INode) pushUp() {
	// 意味着左右子树有一个树的区间有预定，这个树就不能再预定了
	in.val = max(in.left.val, in.right.val)
}

func (in *INode) pushDown(leftNum, rightNum int) {
	if in.left == nil {
		in.left = &INode{}
	}
	if in.right == nil {
		in.right = &INode{}
	}

	if in.lazy == 0 {
		return
	}

	in.left.val = in.left.val + leftNum*in.lazy
	in.right.val = in.right.val + rightNum*in.lazy

	in.left.lazy += in.lazy
	in.right.lazy += in.lazy

	in.lazy = 0
}

type MyCalendar struct {
	*INode
}

func Constructor() MyCalendar {
	return MyCalendar{INode: &INode{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	q := this.query(0, N, start, end-1)
	if q > 0 {
		return false
	}
	this.update(0, N, start, end-1, 1)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func Test_729(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, true, obj.Book(10, 20))
	//assert.Equal(t, false, obj.Book(15, 25))
	assert.Equal(t, true, obj.Book(20, 30))
}
