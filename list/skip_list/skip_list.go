package skip_list

import (
	"fmt"
	"math/rand"
	"time"
)

type SkipList struct {
	maxLen int
	head   *node
}

type node struct {
	nexts    []*node // 这个地方是精髓，它的长度意味着下一个节点的层数
	key, val int
}

type Option func(list *SkipList)

func WithMaxLenOption(max int) Option {
	return func(sl *SkipList) {
		sl.maxLen = max
	}
}

func New(opts ...Option) *SkipList {
	s := &SkipList{
		head: &node{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// 随机获取插入节点的level
// 插入1层的概率是1/2，每增加一层的概率变为 (1/2)^(...+1)
func roll() int {
	rand.Seed(time.Now().UnixNano())
	var level int
	for rand.Intn(2) == 0 {
		level++
	}
	return level
}

func (s *SkipList) Put(key, val int) {
	level := roll()
	if level > s.maxLen && s.maxLen != 0 {
		level = s.maxLen
	}

	// 补齐head的高度, 这里注意要减一，因为是level是从0开始的
	for len(s.head.nexts)-1 < level {
		s.head.nexts = append(s.head.nexts, nil)
	}
	// new一个节点
	newNode := &node{
		nexts: make([]*node, level+1),
		key:   key,
		val:   val,
	}

	// 对每一层进行遍历
	move := s.head
	for curl := level; curl >= 0; curl-- {
		for move.nexts[curl] != nil && move.nexts[curl].key < key {
			move = move.nexts[curl]
		}
		// 插入新节点
		newNode.nexts[curl] = move.nexts[curl]
		move.nexts[curl] = newNode
	}
}

func (s *SkipList) Get(key int) (int, bool) {
	if n := s.search(key); n != nil {
		return n.val, true
	}
	return -1, false
}

func (s *SkipList) search(key int) *node {
	move := s.head
	for curl := len(s.head.nexts) - 1; curl >= 0; curl-- {
		for move.nexts[curl] != nil && move.nexts[curl].key < key {
			move = move.nexts[curl]
		}
		if move.nexts[curl] != nil && move.nexts[curl].key == key {
			return move.nexts[curl]
		}
	}
	return nil
}

func (s *SkipList) Del(key int) {
	move := s.head
	for curl := len(s.head.nexts) - 1; curl >= 0; curl-- {
		for move.nexts[curl] != nil && move.nexts[curl].key < key {
			move = move.nexts[curl]
		}
		if move.nexts[curl] == nil || move.nexts[curl].key != key {
			continue
		}
		move.nexts[curl] = move.nexts[curl].nexts[curl]
	}

	// 防止这种情况出现，删除19之后 => Level 4: nil
	//Level 4: (19, 19) -> nil
	//Level 3: (19, 19) -> (26, 26) -> nil
	//Level 2: (3, 3) -> (19, 19) -> (23, 23) -> (26, 26) -> (37, 37) -> nil
	//Level 1: (3, 3) -> (7, 7) -> (19, 19) -> (23, 23) -> (26, 26) -> (37, 37) -> nil
	//Level 0: (3, 3) -> (7, 7) -> (19, 19) -> (21, 21) -> (23, 23) -> (26, 26) -> (37, 37) -> nil
	//
	//Level 4: nil
	//Level 3: (26, 26) -> nil
	//Level 2: (3, 3) -> (23, 23) -> (26, 26) -> (37, 37) -> nil
	//Level 1: (3, 3) -> (7, 7) -> (23, 23) -> (26, 26) -> (37, 37) -> nil
	//Level 0: (3, 3) -> (7, 7) -> (21, 21) -> (23, 23) -> (26, 26) -> (37, 37) -> nil

	nilCnt := 0
	for curl := len(s.head.nexts) - 1; curl >= 0; curl-- {
		if s.head.nexts[curl] == nil {
			nilCnt++
		}
	}
	s.head.nexts = s.head.nexts[:len(s.head.nexts)-nilCnt]
}

// Range 寻找[start, end]闭区间的节点
func (s *SkipList) Range(start, end int) [][2]int {
	res := make([][2]int, 0, len(s.head.nexts))
	ceilingNode := s.ceiling(start)
	if ceilingNode == nil {
		return res
	}
	for cur := ceilingNode; cur != nil && cur.key <= end; cur = cur.nexts[0] {
		res = append(res, [2]int{cur.key, cur.val})
	}
	return res
}

func (s *SkipList) isNilNode(n *node) bool {
	return n.key == 0
}

// ceil 返回key >= target, 且最接近target的节点
func (s *SkipList) ceiling(target int) *node {
	move := s.head
	for curl := len(s.head.nexts) - 1; curl >= 0; curl-- {
		for move.nexts[curl] != nil && move.nexts[curl].key < target {
			move = move.nexts[curl]
		}
		if move.nexts[curl] != nil && move.nexts[curl].key == target {
			return move.nexts[curl]
		}
	}
	// 这里是与floor()的区别
	return move.nexts[0]
}

// floor 返回key <= target, 且最接近target的节点
func (s *SkipList) floor(target int) *node {
	move := s.head
	for curl := len(s.head.nexts) - 1; curl >= 0; curl-- {
		for move.nexts[curl] != nil && move.nexts[curl].key < target {
			move = move.nexts[curl]
		}
		if move.nexts[curl] != nil && move.nexts[curl].key == target {
			return move.nexts[curl]
		}
	}

	// 由于head 初始化的时候是&node{},所以当target比最小值还小时，返回的是对应类型的空值而不是nil
	if s.isNilNode(move) {
		return nil
	}

	// 这里是与ceiling()的区别
	return move
}

func (s *SkipList) Print() {
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		current := s.head.nexts[level]
		fmt.Printf("Level %d: ", level)
		for current != nil {
			fmt.Printf("(%d, %d) -> ", current.key, current.val)
			current = current.nexts[level]
		}
		fmt.Println("nil")
	}
}
