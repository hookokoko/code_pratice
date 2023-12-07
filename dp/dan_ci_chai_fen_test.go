package dp

import (
	"fmt"
	"testing"
)

func wordBreakWrong(words []string, s string) bool {
	dp := make([]bool, len(s))
	dp[0] = true
	for _, word := range words {
		for j := len(word); j < len(s); j++ {
			subStr := s[j-len(word) : j]
			//fmt.Printf("%s----%s\n", subStr, word)
			if word == subStr && dp[j-len(word)] {
				dp[j] = true
			}
			//fmt.Printf("word[%s], j[%d], %v\n", word, j, dp)
		}
	}
	return dp[len(s)]
}

func wordBreak(words []string, s string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		fmt.Println("=====", i)
		for _, word := range words {
			if i-len(word) >= 0 {
				fmt.Printf("%s----%s\n", s[i-len(word):i], word)
				if dp[i-len(word)] && word == s[i-len(word):i] {
					dp[i] = true
				}
				fmt.Printf("%v\n", dp)
			}
		}
	}
	return dp[len(s)]
}

func TestWordBreak(t *testing.T) {
	// 5, 3
	words := []string{"apple", "pen"}
	// 13
	s := "applepenapple"
	res := wordBreak(words, s)
	println(res)
}
