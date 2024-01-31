package erfen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		r1 := searchKth(nums1, nums2, l/2)
		r2 := searchKth(nums1, nums2, l/2+1)
		return (r1 + r2) / 2
	}
	return searchKth(nums1, nums2, l/2+1)
}

func searchKth(nums1 []int, nums2 []int, k int) float64 {
	var res int
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			res = nums2[idx2+k-1]
			break
		}
		if idx2 == len(nums2) {
			res = nums1[idx1+k-1]
			break
		}
		if k == 1 {
			res = min(nums1[idx1], nums2[idx2])
			break
		}
		half := k / 2
		// 获取分界点的前一个的元素，如果idx+half超过了nums的长度，则取nums的长度
		newIdx1 := min(len(nums1), idx1+half) - 1
		newIdx2 := min(len(nums2), idx2+half) - 1
		// 比较分界点前一个元素的大小
		if nums1[newIdx1] < nums2[newIdx2] {
			// 新的k值是，原来的k值减掉已经比较过的元素个数
			k = k - (newIdx1 - idx1 + 1)
			idx1 = newIdx1 + 1
		} else {
			k = k - (newIdx2 - idx2 + 1)
			idx2 = newIdx2 + 1
		}
	}
	return float64(res)
}

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, float64(4), findMedianSortedArrays([]int{1}, []int{3, 4, 5, 6}))
}
