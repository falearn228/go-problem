// Last updated: 11/7/2025, 2:48:46 PM
package main

type LRUCache struct {
	cap        int
	cache      map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val   int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{key: 0, val: 0}
	tail := &Node{key: 0, val: 0}

	head.next = tail
	tail.prev = head

	return LRUCache{
		cap:   capacity,
		head:  head,
		tail:  tail,
		cache: make(map[int]*Node),
	}
}

func (this *LRUCache) removeNode(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.removeNode(value)
	this.addToHead(value)

	return value.val
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.removeNode(node)
		this.addToHead(node)
		return
	}

	if len(this.cache) == this.cap {
		lruNode := this.tail.prev
		this.removeNode(lruNode)
		delete(this.cache, lruNode.key)
	}

	newNode := &Node{key: key, val: value}
	this.addToHead(newNode)
	this.cache[key] = newNode
}
