package hash

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0, len(strs))
	m := make(map[string][]string, len(strs))
	for _, str := range strs {
		mCnt := [26]rune{}
		for i := 0; i < len(str); i++ {
			mCnt[str[i]-'a']++
		}

		s := strings.Builder{}
		for i := 0; i < 26; i++ {
			for mCnt[i] != 0 {
				s.WriteRune(rune(i) + 'a')
				mCnt[i]--
			}
		}
		// fmt.Println(mCnt, "  ", s.String())
		m[s.String()] = append(m[s.String()], str)
	}
	for _, val := range m {
		res = append(res, val)
	}
	return res
}

func Test_groupAnagrams(t *testing.T) {
	assert.Equal(t, [][]string{{"ddddddddddg"}, {"dgggggggggg"}}, groupAnagrams([]string{"ddddddddddg", "dgggggggggg"}))
}
