package leetcode

type LRUCache struct {
	set        map[int]*DLinkedNode
	head, tail *DLinkedNode
	cap, len   int
}

type DLinkedNode struct {
	prev, next *DLinkedNode
	key        int
	val        int
}

func Constructor(capacity int) LRUCache {
	head := &DLinkedNode{}
	tail := &DLinkedNode{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		set:  make(map[int]*DLinkedNode, capacity+1),
		cap:  capacity,
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.set[key]
	if ok {
		this.moveToHead(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.set[key]
	if ok {
		node.val = value
		this.moveToHead(node)
	} else {
		node = &DLinkedNode{
			prev: nil,
			next: nil,
			key:  key,
			val:  value,
		}
		this.set[key] = node
		this.addToHead(node)
		this.len++
		for this.len > this.cap {
			this.removeTail()
			delete(this.set, node.key)
			this.len--
		}
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.next.prev = node
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) removeTail() {
	node := this.tail.prev
	this.removeNode(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
