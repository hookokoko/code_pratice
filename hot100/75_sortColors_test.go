package hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1, 2, 1, 0, 0, 1, 2
/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
*/
func sortColors(nums []int) {
	p0 := 0             // [0, p1) 用来表示0的区间
	p2 := len(nums) - 1 // [p1, p2) 用来表示1的区间
	// [p2, len(nums)) 用来表示2的区间

	idx := 0 // idx需要自己去控制
	for idx <= p2 {
		if nums[idx] == 0 {
			// 交换位于p0和i值
			swap(nums, p0, idx)
			p0++
			idx++
		} else if nums[idx] == 1 {
			idx++
		} else if nums[idx] == 2 {
			swap(nums, idx, p2)
			p2--
			// 这里不需要idx--, 为什么要这样？
			// nums[idx]=0：此时将其与位置 p0 进行互换（记住 p0 为下一个待填入 0 的位置，同时 [p0,idx−1] 为 1 的区间），
			//		本质是将 nums[idx]nums[idx]nums[idx] 的 0 和 nums[p0] 的 1 进行互换，因此互换后将 l 和 idx 进行右移；
			// nums[idx]=1 nums[idx] = 1 nums[idx]=1：
			//		由于 [p0,idx−1] 本身就是 1 的区间，直接将 idx 进行右移即可；
			// nums[idx]=2 nums[idx] = 2 nums[idx]=2：
			//		此时将其与位置 p2 进行互换（p2 为下一个待填入 2 的位置，但 [idx,p2] 为未处理区间），
			//		也就是我们互换后，只能明确换到位置 nums[p2] 的位置为 2，可以对 p2 进行左移，
			//		但不确定新 nums[idx] 为何值，因此保持 idx 不变再入循环判断。
		}
	}
}

// 顺着照刚刚的思路，有一个更容易理解的算法
func sortColors_easy(nums []int) {
	p0 := 0             // [0, p1) 用来表示0的区间
	p2 := len(nums) - 1 // [p1, p2) 用来表示1的区间
	// [p2, len(nums)) 用来表示2的区间
	for i := 0; i <= p2; i++ {
		// 也就是说这里, 如果i这里的值为2，就需要一直和p2交换，知道新的i的值不为2
		for nums[i] == 2 && i < p2 {
			swap(nums, i, p2)
			p2--
		}
		if nums[i] == 0 {
			swap(nums, i, p0)
			p0++
		}
	}
}

func swap(nums []int, x, y int) {
	if x < 0 || x > len(nums) || y < 0 || y > len(nums) {
		return
	}
	nums[x], nums[y] = nums[y], nums[x]
}

func Test_swap(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		wantNums []int
		i        int
		j        int
	}{
		{name: "1", nums: []int{1, 2, 3, 4, 5, 6, 7}, wantNums: []int{1, 2, 6, 4, 5, 3, 7}, i: 2, j: 5},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.nums, tt.i, tt.j)
			assert.Equal(t, tt.nums, tt.wantNums)
		})
	}
}

func Test_sortColors(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		wantNums []int
	}{
		{name: "1", nums: []int{1, 2, 1, 0, 0, 1, 2}, wantNums: []int{0, 0, 1, 1, 1, 2, 2}},
		{name: "2", nums: []int{2, 0, 2, 1, 1, 0}, wantNums: []int{0, 0, 1, 1, 2, 2}},
		{name: "3", nums: []int{2}, wantNums: []int{2}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			sortColors_easy(tt.nums)
			assert.Equal(t, tt.wantNums, tt.nums)
		})
	}
}
