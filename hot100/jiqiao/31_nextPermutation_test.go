package jiqiao

import "testing"

// 找到序列中下一个 **尽可能小的** 更大的数
// 2 3 1 ==> 3 1 2
// 宗旨是交换前面的小数和后面的大数，才能是序列值更大。但同时要保证变大的幅度最小
func nextPermutation(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	for i := len(nums) - 1; i > 0; i-- {
		// 先找到第一个递增的连续数对
		if nums[i-1] < nums[i] {
			// 从后往前找第一个比数对左边值大的元素，并与之交换
			// 然后逆序这个左值之后的部分
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					reverse(nums[i:])
					return
				}
			}
		}
	}
	reverse(nums)
}

func nextPermutationBetter(nums []int) {
	// 找到第一个递增的连续数对，i为递增对的右值
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}

	// 从后往前找第一个比数对左边值大的元素，并与之交换
	// 然后逆序这个左值之后的部分
	if i > 0 {
		for j := len(nums) - 1; j >= i; j-- {
			if nums[j] > nums[i-1] {
				nums[i-1], nums[j] = nums[j], nums[i-1]
				break
			}
		}
	} else {
		// 无需再交换了
	}

	reverse(nums[i:])
}

func reverse(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func Test_nextPermutation(t *testing.T) {
	nums := []int{2, 3, 1}
	//nextPermutation(nums)
	nextPermutationBetter(nums)

	t.Log(nums)
}
