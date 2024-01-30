package erfen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchRange(nums []int, target int) []int {
	low := binarySearch(nums, target, true)
	high := binarySearch(nums, target, false) - 1
	return []int{low, high}
}

// 0, 1, 2, 3, 4, 5, 6, 7, 8
// 5, 7, 7, 7, 7, 7, 8, 8, 10

/* 右边界
[0, 8], mid=4, nums[4]=7, left=mid+1=5
[5, 8], mid=6, nums[6]=8, right=mid-1=5
[5, 5], mid=5, nums[5]=7, left=mid+1=6
*/
/* 左边界
[0, 8], mid=4, nums[4]=7, right=mid-1=3
[0, 3], mid=1, nums[1]=7, right=mid-1=0
[0, 0], mid=0, nums[0]=5, left=mid+1=1
*/

func binarySearch(nums []int, target int, low bool) int {
	left := 0
	right := len(nums) - 1
	ans := len(nums)
	for left <= right {
		mid := left + (right-left)>>1
		if !low {
			// 意味着相等的话，left一直会加1，直到到右边界
			if nums[mid] > target {
				right = mid - 1
				ans = mid
			} else {
				left = mid + 1
			}
		} else {
			// 意味着相等的话，right一直会减1，直到左边界
			if nums[mid] >= target {
				right = mid - 1
				ans = mid
			} else {
				left = mid + 1
			}
		}
	}
	return ans
}

func TestSearchRange(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			name:   "2",
			nums:   []int{5, 7, 7, 7, 8, 8, 10},
			target: 7,
			want:   []int{1, 3},
		},
		{
			name:   "3",
			nums:   []int{5, 7, 7, 7, 7, 7, 8, 8, 10},
			target: 7,
			want:   []int{1, 5},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchRange(tt.nums, tt.target))
		})
	}
}
