package sub_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var q []int

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	q = make([]int, 0, len(nums))
	for i := 0; i < k; i++ {
		push(nums, i)
	}
	res = append(res, nums[q[0]])
	// i是从第k个位置开始，也就是k-1
	for i := k; i < len(nums); i++ {
		//fmt.Printf("i:%d-------q:%v------res:%v\n", i, q, res)
		push(nums, i)
		// 可以不加=吗？这里肯定是不可以的，因为已经多push了一个
		for len(q) > 0 && i-q[0] >= k {
			q = q[1:]
		}
		res = append(res, nums[q[0]])
	}
	return res
}

func push(nums []int, i int) {
	// 可以不加=吗？可以，因为队列里面的判断是通过索引的距离
	for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
		q = q[:len(q)-1]
	}
	q = append(q, i)
}

func Test_maxSlidingWindow(t *testing.T) {
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	assert.Equal(t, []int{7, 4}, maxSlidingWindow([]int{7, 2, 4}, 2))
	assert.Equal(t, []int{3, 3, 2, 5}, maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
