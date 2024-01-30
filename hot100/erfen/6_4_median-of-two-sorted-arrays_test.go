package erfen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return searchKth(nums1, nums2, 2)
}

func searchKth(nums1 []int, nums2 []int, k int) float64 {
	idx1, idx2 := 0, 0
	for {
		half := k / 2
		newIdx1 := min(len(nums1), idx1+half) - 1
		newIdx2 := min(len(nums2), idx2+half) - 1
		if nums1[newIdx1] < nums2[newIdx2] {
			k = k - (newIdx1 - idx1 + 1)
			idx1 = newIdx1 + 1
		} else {
			k = k - (newIdx2 - idx2 + 1)
			idx2 = newIdx2 + 1
		}
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, 2, findMedianSortedArrays([]int{1, 3}, []int{2}))
	//assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	//assert.Equal(t, 0, findMedianSortedArrays([]int{}, []int{}))
	//assert.Equal(t, 0, findMedianSortedArrays([]int{}, []int{}))
}
