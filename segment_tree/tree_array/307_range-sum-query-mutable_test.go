package tree_array

import (
	"testing"
)

var sum []int

func sumRange(arr []int, low, high int) int {
	return count(arr, high) - count(arr, low)
}

func count(arr []int, idx int) int {
	var res int
	for idx != 0 {
		res = res + arr[idx]
		idx -= lowBit(idx)
	}
	return res
}

func update(arr []int, idx, val int) {
	for idx < len(arr) {
		arr[idx] += val
		idx += lowBit(idx)
	}
}

func Test_SumRange(t *testing.T) {
	arr := []int{1, 3, 5}
	sum = make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		update(sum, i+1, arr[i])
	}
	t.Log(sum)
	t.Log(count(sum, 1))
	t.Log(count(sum, 2))
}
