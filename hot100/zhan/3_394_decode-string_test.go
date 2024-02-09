package zhan

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// 两个栈，一个保存数字，一个保存字符串
// 写个例子，对着想：3[a2[cd]11[e]]
func decodeString(s string) string {
	var nums = make([]int, 0)
	var strs = make([]string, 0)

	var res string
	var num int
	for _, item := range s {
		itemStr := string(item)
		if itemStr >= "0" && itemStr <= "9" {
			num = num*10 + toInt(itemStr)
		} else if itemStr >= "a" && itemStr <= "z" {
			res += itemStr
		} else if itemStr == "[" {
			nums = append(nums, num)
			num = 0
			strs = append(strs, res)
			res = ""
		} else if itemStr == "]" {
			for i := 0; i < nums[len(nums)-1]; i++ {
				strs[len(strs)-1] += res
			}
			res = strs[len(strs)-1]
			strs = strs[:len(strs)-1]
			nums = nums[:len(nums)-1]
		} else {
			return "illegal"
		}
	}
	return res
}

func toInt(s string) int {
	ss, _ := strconv.Atoi(s)
	return ss
}

func TestDecodeString(t *testing.T) {
	assert.Equal(t, "acdcdeeeeeeeeeeeacdcdeeeeeeeeeeeacdcdeeeeeeeeeee", decodeString("3[a2[cd]11[e]]"))
}
