package skip_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_roll(t *testing.T) {
	c := make(map[int]int)
	for i := 0; i < 10000; i++ {
		c[roll()]++
	}
	fmt.Println(c)
}

func Test_Put(t *testing.T) {
	s := _new()
	s.Print()
}

func Test_Get(t *testing.T) {
	s := _new()
	val, ok := s.Get(37)
	fmt.Println(val)
	fmt.Println(ok)
}

func Test_Del(t *testing.T) {
	s := _new()
	s.Print()
	s.Del(19)
	fmt.Printf("\n")
	s.Print()
}

func Test_ceiling(t *testing.T) {
	s := _new()
	n := s.ceiling(40)
	assert.Nil(t, n)

	n = s.ceiling(20)
	fmt.Printf("key: %d, val: %d\n", n.key, n.val)
}

func Test_floor(t *testing.T) {
	s := _new()
	n := s.floor(1)
	fmt.Printf("key: %d, val: %d\n", n.key, n.val)
}

func Test_Range(t *testing.T) {
	s := _new()
	res := s.Range(3, 9)
	fmt.Printf("%+v\n", res)
}

func _new() *SkipList {
	s := New()
	s.Put(3, 3)
	s.Put(37, 37)
	s.Put(19, 19)
	s.Put(7, 7)
	s.Put(21, 21)
	s.Put(23, 23)
	s.Put(26, 26)
	return s
}

func Test_other(t *testing.T) {
	s := make([]*node, 1)
	fmt.Printf("%+v\n", s)
}
