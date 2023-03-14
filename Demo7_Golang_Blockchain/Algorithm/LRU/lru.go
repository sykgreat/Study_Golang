package LRU

import (
	"golang.org/x/exp/constraints"
)

type DLinkNode[K constraints.Ordered, V any] struct {
	key        K
	val        V
	prev, next *DLinkNode[K, V]
}

func NewDLinkNode[K constraints.Ordered, V any]() *DLinkNode[K, V] {
	return &DLinkNode[K, V]{}
}

type LRUCache[K constraints.Ordered, V any] struct {
	size, capacity int
	cache          map[K]*DLinkNode[K, V]
	head, tail     *DLinkNode[K, V]
}

func Constructor[K constraints.Ordered, V any](capacity int) LRUCache[K, V] {
	lru := LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*DLinkNode[K, V], capacity),
		head:     NewDLinkNode[K, V](),
		tail:     NewDLinkNode[K, V](),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (node *DLinkNode[K, V]) removeNode() {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCache[K, V]) removeTail() *DLinkNode[K, V] {
	node := lru.tail.prev
	node.removeNode()
	return node
}

func (lru *LRUCache[K, V]) addToHead(node *DLinkNode[K, V]) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCache[K, V]) moveToHead(node *DLinkNode[K, V]) {
	node.removeNode()
	lru.addToHead(node)
}

func (lru *LRUCache[K, V]) Get(key K) (v V) {
	if node, ok := lru.cache[key]; ok {
		lru.moveToHead(node)
		v = node.val
	}
	return v
}

func (lru *LRUCache[K, V]) Put(key K, value V) {
	if node, ok := lru.cache[key]; ok {
		node.val = value
		lru.moveToHead(node)
	} else {
		node := &DLinkNode[K, V]{key: key, val: value}
		lru.cache[key] = node
		lru.addToHead(node)
		lru.size++
		if lru.size > lru.capacity {
			removed := lru.removeTail()
			delete(lru.cache, removed.key)
			lru.size--
		}
	}
}
