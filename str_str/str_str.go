package str_str

import (
	"code_pratice/snips/algorithm"
	"fmt"
)

/*
 	没有必要每一个算法都建一个文件
	字符串相关的算法暂时都放在这
	必须加好注释
*/

// ReverseWords 单词翻转，难点在于其中的多余空格的删除
func ReverseWords(s string) string {
	sb := []byte(s)
	algorithm.RemoveExtraSpace(&sb)
	fmt.Println(string(sb))
	reverse(sb)
	fmt.Println(string(sb))
	recordIdx := 0
	for i := 0; i <= len(sb); i++ {
		if i == len(sb) || sb[i] == '_' {
			reverse(sb[recordIdx:i])
			recordIdx = i + 1
		}
	}
	return string(sb)
}

func reverseWords(s string) string {
	b := []byte(s)

	// 移除前面、中间、后面存在的多余空格
	slow := 0
	for i := 0; i < len(b); i++ {
		if b[i] != '_' {
			if slow != 0 {
				b[slow] = '_'
				slow++
			}
			for i < len(b) && b[i] != '_' { // 复制逻辑
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	b = b[0:slow]

	// 翻转整个字符串
	reverse(b)
	// 翻转每个单词
	last := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == '_' {
			reverse(b[last:i])
			last = i + 1
		}
	}
	return string(b)
}

func reverse(sb []byte) {
	left := 0
	right := len(sb) - 1
	for left < right {
		sb[left], sb[right] = sb[right], sb[left]
		left++
		right--
	}
}
