package array

import "math"

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。子数组是数组中的一个连续部分。
*/

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := -math.MaxInt
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
