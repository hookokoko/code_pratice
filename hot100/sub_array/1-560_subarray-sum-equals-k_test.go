package sub_array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func subarraySum(nums []int, k int) int {
	m := make(map[int]int, len(nums)+1)
	m[0] = 1
	pre := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			cnt += m[pre-k]
			fmt.Printf("		nums[%d]=%d, cnt=%d\n", i, nums[i], cnt)
		}
		m[pre]++
		fmt.Println(m)
	}
	return cnt
}

func subarraySum1(nums []int, k int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if sum(nums, i, j) == k { // 通用需要for遍历，时间复杂度O(n)
				cnt++
			}
		}
	}
	return cnt
}
func sum(nums []int, start, end int) int {
	res := 0
	for i := start; i < end; i++ {
		res += nums[i]
	}
	return res
}

func subarraySum2(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		pre := 0
		for end := start; end >= 0; end-- {
			pre += nums[end]
			if pre == k {
				count++
			}
		}
	}
	return count
}

func Test_subarraySum(t *testing.T) {
	assert.Equal(t, 4, subarraySum([]int{3, 4, 7, 2, -3, 1, 4, 2}, 7))
	//assert.Equal(t, 2, subarraySum([]int{1, 1, 1}, 2))

}
