package huadongchuangkou

import (
	"fmt"
	"testing"
)

/* 438. 找出字符串中所有字母异位词
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
*/

func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	res := make([]int, 0, len(s))
	pCnt := [26]rune{}
	sCnt := [26]rune{}
	for _, pp := range p {
		pCnt[pp-'a']++
	}

	for right < len(s) {
		sCnt[s[right]-'a']++
		if right-left+1 == len(p) {
			if pCnt == sCnt {
				res = append(res, left)
			}
			sCnt[s[left]-'a']--
			left++
		}
		right++
	}
	return res
}

func Test_11(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
