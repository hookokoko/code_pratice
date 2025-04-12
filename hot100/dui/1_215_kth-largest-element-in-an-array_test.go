package dui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findKthLargest(nums []int, k int) int {
	//return quickselect(nums, 0, len(nums)-1, len(nums)-k)
	return quickselect1(nums, 0, len(nums)-1, k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i = i + 1; nums[i] < partition; i++ {
		}
		for j = j - 1; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	}
	return quickselect(nums, j+1, r, k)
}

func quickselect1(nums []int, l, r, k int) int {
	partition := nums[l]
	i := l
	j := r
	for i < j {
		for i < j && nums[i] >= partition {
			i++
		}
		for i < j && nums[j] <= partition {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i], nums[l] = nums[l], nums[i]

	if i == k-1 {
		return nums[i]
	} else if i < k-1 {
		return quickselect1(nums, i+1, r, k)
	}
	return quickselect1(nums, l, j-1, k)

	//if j == k {
	//	return nums[k]
	//} else if j < k {
	//	return quickselect1(nums, j+1, r, k)
	//}
	//return quickselect1(nums, l, j-1, k)
}

func Test_FindKthLargest(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "case1",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "case2",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "case3",
			nums: []int{1, 2, 2},
			k:    3,
			want: 2,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
