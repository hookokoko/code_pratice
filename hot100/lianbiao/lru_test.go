package lianbiao

import (
	"fmt"
	"testing"
)

type LRUCache struct {
	data map[int]*DLinkListNode
	cap  int
	size int
	head *DLinkListNode
	tail *DLinkListNode
}

func Constructor(capacity int) LRUCache {
	dummyHead := &DLinkListNode{}
	dummyTail := &DLinkListNode{}
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead
	return LRUCache{
		data: make(map[int]*DLinkListNode, capacity),
		cap:  capacity,
		head: dummyHead,
		tail: dummyTail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.data[key]
	if ok {
		node.Val = value
		this.moveToHead(node)
		return
	}
	newNode := &DLinkListNode{Val: value, Key: key}
	this.size += 1
	if this.size > this.cap {
		this.popTail()
		this.size -= 1
	}
	this.data[key] = newNode
	this.addToHead(newNode)
	return
}

type DLinkListNode struct {
	Prev *DLinkListNode
	Next *DLinkListNode
	Val  int
	Key  int // key不能省
}

func (this LRUCache) moveToHead(node *DLinkListNode) {
	// 先删除node
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	// 再将node add到head
	this.addToHead(node)
}

func (this LRUCache) popTail() {
	toDel := this.tail.Prev
	toDel.Prev.Next = this.tail
	this.tail.Prev = toDel.Prev
	delete(this.data, toDel.Key)
}

func (this LRUCache) addToHead(node *DLinkListNode) {
	// 先修改node的指向。因为这样不影响原来的节点
	node.Next = this.head.Next
	node.Prev = this.head
	// 再修改this.head的下一个节点的指向，下面两步顺序不能变
	// 实际上this.head是一个虚拟节点，改变的是其后的节点
	this.head.Next.Prev = node
	this.head.Next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Test_1(t *testing.T) {
	obj := Constructor(2)
	fmt.Println(obj.Get(1))

	obj.Put(1, 0)
	obj.Put(2, 2)
	Print(obj)
	obj.Get(1)
	Print(obj)

	obj.Put(3, 3)
	Print(obj)
	obj.Get(2)
	Print(obj)
	obj.Put(4, 4)
	Print(obj)

	fmt.Println("-----", obj.Get(1))
	Print(obj)
	obj.Get(3)
	Print(obj)
	obj.Get(4)
	Print(obj)
}

func Print(obj LRUCache) {
	var res string
	for node := obj.head.Next; node != obj.tail; node = node.Next {
		res += fmt.Sprintf("%d ->", node.Val)
	}
	fmt.Println(res)
}
