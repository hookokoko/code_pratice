package segment_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Node struct {
	left  *Node
	right *Node
	val   int
	lazy  int
}

func New(arr []int, start, end int) *Node {
	node := &Node{}
	node.build(arr, start, end)
	return node
}

func (n *Node) build(arr []int, start, end int) {
	if start == end {
		n.val = arr[start]
		return
	}
	n.left = &Node{}
	n.right = &Node{}
	mid := start + (end-start)>>1
	n.left.build(arr, start, mid)
	n.right.build(arr, mid+1, end)
	n.pushUp()
	return
}

func (n *Node) query(start, end, left, right int) int {
	// 终止条件的说明
	// 实际上到这里，要么是叶子结点了，要么是[start, end]是[left, right]的子区间了。子区间的特点是，要么start==left，要么end==right
	// 比如在[0,4]找[1,2]，最后就是[1,1]中找[1,2], [2,2]中找[1,2]这种了；
	// 再比如,[0,4]中找[2,4],最后就是[2,2]中找[2,4]，[3,4]中找[2,4]。
	// [start, end]中找[left, right]
	if left <= start && right >= end {
		return n.val
	}
	ans := 0
	mid := start + (end-start)>>1
	n.pushDown(mid-start+1, end-mid)
	if left <= mid {
		// 变化的只有start、end
		ans += n.left.query(start, mid, left, right)
	}
	// 不能是else，因为mid可能在 left 和 right 之间
	if right > mid {
		ans += n.right.query(mid+1, end, left, right)
	}
	return ans
}

func (n *Node) add(start, end, left, right, val int) {
	if left <= start && right >= end {
		n.val += (end - start + 1) * val
		n.lazy += val
		return
	}
	mid := start + (end-start)>>1
	n.pushDown(mid-start+1, end-mid)
	// 无需对left、right是否为nil进行判断，因为pushDown会做
	if left <= mid {
		n.left.add(start, mid, left, right, val)
	}
	// 不能是else，因为mid可能在left和right之间
	if right > mid {
		n.right.add(mid+1, end, left, right, val)
	}
	n.pushUp()
}

func (n *Node) pushUp() {
	n.val = n.left.val + n.right.val
}

func (n *Node) pushDown(leftNum, rightNum int) {
	if n.left == nil {
		n.left = &Node{}
	}
	if n.right == nil {
		n.right = &Node{}
	}
	if n.lazy == 0 {
		return
	}

	n.left.val += n.lazy * leftNum
	n.right.val += n.lazy * rightNum

	n.left.lazy += n.lazy
	n.right.lazy += n.lazy
	n.lazy = 0
}

func Test_1(t *testing.T) {
	tree := New([]int{1, 2, 3, 4, 5}, 0, 4)
	//assert.Equal(t, 15, tree.query(0, 4, 0, 4))
	//assert.Equal(t, 5, tree.query(0, 4, 1, 2))
	//assert.Equal(t, 10, tree.query(0, 4, 0, 3))
	//assert.Equal(t, 9, tree.query(0, 4, 1, 3))
	tree.add(0, 4, 1, 2, 1)
	assert.Equal(t, 7, tree.query(0, 4, 1, 2))
}
