package erfen

import (
	"fmt"
	"testing"
)

/*
搜索旋转排序数组,利用两次二分
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	// 第一次二分找到旋转点
	for left <= right {
		mid := left + (right-left)>>1
		// 原数组数字各不相同，实际取不到=
		if nums[mid] >= nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//fmt.Println(left)
	// 更新第二次二分的范围
	if target >= nums[0] {
		right = left - 1
		left = 0
	} else {
		right = len(nums) - 1
	}
	//fmt.Println(left, right)
	// 普通的二分查找了
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		} else {
			left = mid + 1
		}
		//fmt.Printf("%d-%d-%d\n", left, right, mid)
	}
	return -1
}

func Test_1(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
}
