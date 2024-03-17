package sub_string

import (
	"math"
	"testing"
)

func minWindow(s string, t string) string {
	tCnt := make(map[string]int)
	tmpCnt := make(map[string]int)
	// 目标串的字符计数
	for i := 0; i < len(t); i++ {
		tCnt[string(t[i])]++
	}
	check := func() bool {
		//这个坑会导致超时，因为每次比较都会遍历t数组
		//for i := 0; i < len(t); i++ {
		//	if tmpCnt[string(t[i])] < tCnt[string(t[i])] {
		//		return false
		//	}
		//}
		for k, v := range tCnt {
			if tmpCnt[k] < v {
				return false
			}
		}
		return true
	}
	minLen := math.MaxInt
	minL, minR := -1, -1
	l := 0
	r := 0
	for ; r < len(s); r++ {
		// 比较串增加s[r]统计
		tmpCnt[string(s[r])]++
		// 是否满足子串
		for check() && l <= r {
			//更新最小长度并，更新左右指针
			if r-l+1 < minLen {
				minLen = r - l + 1
				minL = l
				minR = r
			}
			// l++之后要在tmpCnt对应的s[l]这里减1
			if _, ok := tmpCnt[string(s[l])]; ok {
				tmpCnt[string(s[l])]--
			}
			l++
		}
	}
	if minL == -1 {
		return ""
	}
	return s[minL : minR+1]
}

func Test111(t *testing.T) {
	//t.Log(minWindow("ADOBECODEBANC", "ABC"))
	//t.Log(minWindow("a", "a"))
}
