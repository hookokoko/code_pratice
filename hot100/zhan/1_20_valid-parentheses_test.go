package zhan

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var parentheses = map[string]string{
	"]": "[",
	"}": "{",
	")": "(",
}

func isValid(s string) bool {
	var stack []int32
	for _, item := range s {
		//fmt.Printf("%c\n", item)
		//fmt.Println(string(item))
		if string(item) == "(" || string(item) == "{" || string(item) == "[" {
			stack = append(stack, item)
		} else if len(stack) > 0 && parentheses[string(item)] == string(stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("([{{}}{}()])"))
	assert.Equal(t, true, isValid("()[]{}"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, false, isValid("("))
}
