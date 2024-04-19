package skip_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
	"time"
)

type Skiplist struct {
	head *Node
}

type Node struct {
	nexts []*Node
	val   int
}

func Constructor() Skiplist {
	head := &Node{
		//nexts: make([]*Node, 0, 32),
	}
	return Skiplist{
		head: head,
	}
}

func randLvl() int {
	rand.Seed(time.Now().UnixNano())
	lvl := 1
	for rand.Intn(100) < 50 {
		lvl += 1
	}
	return lvl
}

func (this *Skiplist) Search(target int) bool {
	move := this.head
	for lvl := len(this.head.nexts) - 1; lvl >= 0; lvl-- {
		for move.nexts[lvl] != nil && move.nexts[lvl].val < target {
			move = move.nexts[lvl]
		}
		if move.nexts[lvl] != nil && move.nexts[lvl].val == target {
			return true
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	lvl := randLvl()
	for len(this.head.nexts) < lvl {
		this.head.nexts = append(this.head.nexts, nil)
	}
	newNode := &Node{
		nexts: make([]*Node, lvl),
		val:   num,
	}
	move := this.head
	for curlvl := lvl - 1; curlvl >= 0; curlvl-- {
		for move.nexts[curlvl] != nil && move.nexts[curlvl].val < num {
			move = move.nexts[curlvl]
		}
		newNode.nexts[curlvl] = move.nexts[curlvl]
		move.nexts[curlvl] = newNode
	}
}

func (this *Skiplist) Erase(num int) bool {
	if this.Search(num) == false {
		return false
	}
	move := this.head
	for lvl := len(this.head.nexts) - 1; lvl >= 0; lvl-- {
		for move.nexts[lvl] != nil && move.nexts[lvl].val < num {
			move = move.nexts[lvl]
		}
		if move.nexts[lvl] != nil && move.nexts[lvl].val == num {
			move.nexts[lvl] = move.nexts[lvl].nexts[lvl]
		}
	}
	nilCnt := 0
	for i := len(this.head.nexts) - 1; i >= 0; i-- {
		if this.head.nexts[i] == nil {
			nilCnt++
		}
	}
	this.head.nexts = this.head.nexts[:len(this.head.nexts)-nilCnt]
	return true
}

func (this *Skiplist) Print() {
	levelRes := make([][]string, 0, 32)
	for level := len(this.head.nexts) - 1; level >= 0; level-- {
		current := this.head.nexts[level]
		//fmt.Printf("Level %d: ", level)
		tmp := make([]string, 0, 32)
		for current != nil {
			//fmt.Printf("%d -> ", current.val)
			tmp = append(tmp, fmt.Sprintf("%d", current.val))
			current = current.nexts[level]
		}
		//fmt.Println("nil")
		tmp = append(tmp, "nil")
		levelRes = append(levelRes, tmp)
	}
	for i := len(levelRes) - 1; i >= 0; i-- {
		diff := len(levelRes[len(levelRes)-1]) - len(levelRes[i])
		emptyStr := make([]string, diff)
		for j := range emptyStr {
			emptyStr[j] = " "
		}
		levelRes[i] = append(emptyStr, levelRes[i]...)
	}
	for i := 0; i < len(levelRes); i++ {
		fmt.Printf("level[%d]: %s\n", len(levelRes)-i-1, strings.Join(levelRes[i], " -> "))
	}
}

func Test_case(t *testing.T) {
	obj := Constructor()
	obj.Add(0)
	obj.Add(5)
	obj.Add(2)
	obj.Add(1)

	obj.Search(0)

	fmt.Println("erase 5", obj.Erase(5))
	fmt.Println("search 2", obj.Erase(2))
	fmt.Println("search 3", obj.Search(3))
	fmt.Println("search 2", obj.Search(2))
}

func Test_case1(t *testing.T) {
	obj := Constructor()

	actions := []string{"add", "add", "add", "add", "add", "erase", "erase", "add", "search", "search", "add", "erase", "search", "add", "add", "add", "erase", "search", "erase", "search", "search", "search", "erase", "erase", "search", "erase", "add", "add", "erase", "add", "search", "search", "search", "search", "search"}
	params := []int{9, 4, 5, 6, 9, 2, 1, 2, 7, 4, 5, 6, 5, 6, 7, 4, 3, 6, 3, 4, 3, 8, 7, 6, 7, 4, 1, 6, 3, 4, 7, 6, 1, 0, 3}
	res := []any{nil, nil, nil, nil, nil, false, false, nil, false, true, nil, true, true, nil, nil, nil, false, false, false, true, false, false, true, true, false, true, nil, nil, false, nil, false, true, true, false, false}

	for i := 0; i < len(actions); i++ {
		switch actions[i] {
		case "add":
			obj.Add(params[i])
		case "erase":
			got := obj.Erase(params[i])
			if res[i] != got {
				fmt.Printf("[%d]: erase error: [%d]\n", i, params[i])
				obj.Print()
			}
		case "search":
			got := obj.Search(params[i])
			if res[i] != got {
				fmt.Printf("[%d]: search error: [%d]\n", i, params[i])
				obj.Print()
			}
		}
	}
	// debug 根据打印结果重新构造数进行debug
}

func Test_case2(t *testing.T) {
	obj := Constructor()

	//level[4]:   -> 4 ------------------------------------> nil
	//level[3]:   -> 4 ------> 5 --------------------------> nil
	//level[2]:   -> 4 -> 4 -> 5 ------> 6 ------> 9 ------> nil
	//level[1]:   -> 4 -> 4 -> 5 ------> 6 -> 7 -> 9 -> 9 -> nil
	//level[0]: 2 -> 4 -> 4 -> 5 -> 5 -> 6 -> 7 -> 9 -> 9 -> nil
	//          1    2    3    4    5    6    7    8    9    10

	c1 := []*Node{{val: 2}}
	c2 := []*Node{{val: 4}, {val: 4}, {val: 4}, {val: 4}, {val: 4}}
	c1[0].nexts = c2
	c3 := []*Node{{val: 4}, {val: 4}, {val: 4}}
	c2[0].nexts = c3
	c2[1].nexts = c3
	c2[2].nexts = c3
	c4 := []*Node{{val: 5}, {val: 5}, {val: 5}, {val: 5}}
	c3[0].nexts = c4
	c3[1].nexts = c4
	c3[2].nexts = c4
	c2[3].nexts = c4
	c5 := []*Node{{val: 5}}
	c4[0].nexts = c5
	c6 := []*Node{{val: 6}, {val: 6}, {val: 6}}
	c5[0].nexts = c6
	c4[1].nexts = c6
	c4[2].nexts = c6
	c7 := []*Node{{val: 7}, {val: 7}}
	c6[0].nexts = c7
	c6[1].nexts = c7
	c8 := []*Node{{val: 9}, {val: 9}, {val: 9}}
	c7[0].nexts = c8
	c7[1].nexts = c8
	c6[2].nexts = c8
	c9 := []*Node{{val: 9}, {val: 9}}
	c8[0].nexts = c9
	c8[1].nexts = c9

	c9[0] = nil
	c9[1] = nil

	c8[2].nexts = nil
	c4[3].nexts = nil
	c2[4].nexts = nil

	obj.head.nexts = c1

	fmt.Println(obj.Search(6))
}

func Test_Case3(t *testing.T) {
	obj := Constructor()

	actions := []string{"add", "add", "add", "add", "add", "erase", "erase", "add", "search", "search", "add", "erase", "search", "add", "add", "add", "erase", "search"}
	params := []int{9, 4, 5, 6, 9, 2, 1, 2, 7, 4, 5, 6, 5, 6, 7, 4, 3, 6}
	res := []any{nil, nil, nil, nil, nil, false, false, nil, false, true, nil, true, true, nil, nil, nil, false, false}

	assert.Equal(t, len(actions), 18)
	assert.Equal(t, len(params), 18)
	assert.Equal(t, len(res), 18)

	for i := 0; i < len(actions); i++ {
		switch actions[i] {
		case "add":
			obj.Add(params[i])
		case "erase":
			got := obj.Erase(params[i])
			if res[i] != got {
				fmt.Printf("[%d]: erase error: [%d]\n", i, params[i])
				obj.Print()
			}
		case "search":
			got := obj.Search(params[i])
			if res[i] != got {
				fmt.Printf("[%d]: search error: [%d]\n", i, params[i])
				obj.Print()
			}
		}
	}
}
