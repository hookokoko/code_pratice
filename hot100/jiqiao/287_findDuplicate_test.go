package jiqiao

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * 快慢指针
 * 要理解的是快慢指针是如何构造，
 * 例如[3,1,3,4,2]，0 -> nums[0]: 3 -> nums[3]: 1 -> nums[1]: 3 -> nums[3]: 4 -> nums[4]: 2 -> nums[2]: 3 ...
 * 再例如[1,3,4,2,2], 0 -> nums[0]: 1 -> nums[1]: 3 -> nums[3]: 2 -> nums[4]: 2 -> nums[2]: 4 -> nums[4]: 2 ...
 * 因此走一步是nums[i], 走两步就是nums[nums[i]]，以此类推
 */
func findDuplicate_1(nums []int) int {
	var slow = 0
	var fast = 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	var pre1 = 0
	var pre2 = slow
	for pre1 != pre2 {
		pre1 = nums[pre1]
		pre2 = nums[pre2]
	}
	return pre1
}

/*
 * 二分查找
 * 二分查找的思路是，从1到n找重复的数，其中len(nums)为n+1
 */

// 这个二分查找的例子，是找<=mid的个数进行判断
func findDuplicate_2_1(nums []int) int {
	// 从1到n找重复的数，其中len(nums)为n+1
	// 此时二分查找是左闭右开区间
	left := 1
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		cnt := 0

		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		// 小于mid的数字的个数“更多”，从左区间继续查找
		if cnt > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 这个二分查找的例子，是找>=mid的个数进行判断
func findDuplicate_2_2(nums []int) int {
	left := 1
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)>>1
		cnt := 0
		for i := 0; i < len(nums); i++ {
			// 这里和下面的cnt > mid的判断是联动的
			if nums[i] >= mid {
				cnt++
			}
		}
		// 核心宗旨是要把mid的值带上，因为mid也可能是重复值。
		// 需要和2-1的例子一起品一品。就按照最简单的例子品：
		// [1,2,3,4,5,6,7,7]这个最普通的例子，稍加改造，让mid就是重复值[1,2,3,4,4,5,6,7]，范围1~7的8个数，第一轮，mid=4。
		// cnt的计算都取"="的条件下，如果是计算<=mid的个数，cnt=5。如果是计算>=mid的个数，cnt也是5；
		// 因此后者要保证4在下一轮时不能被去掉，所以如果是<=mid的cnt，cnt>mid时，right=mid；所以如果是>=mid的cnt，cnt>mid时，left=mid；
		fmt.Printf("left:[%d], mid:[%d], right:[%d], cnt:[%d]\n", left, mid, right, cnt)
		// 总之，因为cnt时＞=mid的值，并且mid可能是重复值，所以left的赋值的时候要包括mid，不能+1了
		if cnt > mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

// 就看这个例子 [2,1,2]
func findDuplicate_3(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i]-1 == i { // 已经排好序
			i++
		} else if nums[nums[i]-1] == nums[i] { // 正确的索引处已经有一个值了
			return nums[i]
		} else {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	return -1
}

// bit位之法，学习学习bit操作，理解不是很难
func findDuplicate_4(nums []int) int {
	n := len(nums)
	ans := 0
	bit_max := 31
	// 右移bit_max位, 相当于除2，向下取整，比如101 >> 1 = 50
	// 看看n是有多少位，因为数字最大就是n-1。比如范围在1~7的8个数，此时bit_max=2
	for ((n - 1) >> bit_max) == 0 {
		bit_max--
	}
	for bit := 0; bit <= bit_max; bit++ {
		x, y := 0, 0
		for i := 0; i < n; i++ {
			if (nums[i] & (1 << bit)) > 0 {
				x++
			}
			if i >= 1 && (i&(1<<bit)) > 0 {
				y++
			}
		}
		// 对二进制的每一位进行一个判断，x、y不相同的情况下，该位就是1，最终计算ans
		if x > y {
			ans |= 1 << bit
		}
	}
	return ans
}

func Test_findDuplicate(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{name: "1", nums: []int{1, 3, 4, 2, 2}, want: 2},
		{name: "2", nums: []int{1, 2, 3, 3, 5, 7, 3, 6}, want: 3},
		{name: "3", nums: []int{1, 2, 4, 4, 5, 7, 3, 6}, want: 4},
		{name: "4", nums: []int{2, 1, 2}, want: 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := findDuplicate_2_2(tt.nums)
			assert.Equal(t, tt.want, res)
		})
	}
}
