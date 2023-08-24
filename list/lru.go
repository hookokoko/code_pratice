package list

type LRUCache struct {
	head     *DLinkedListNode
	tail     *DLinkedListNode
	cache    map[int]*DLinkedListNode
	size     int
	capacity int
}

type DLinkedListNode struct {
	prev *DLinkedListNode
	next *DLinkedListNode
	key  int
	val  int
}

func Constructor(capacity int) LRUCache {

}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}
