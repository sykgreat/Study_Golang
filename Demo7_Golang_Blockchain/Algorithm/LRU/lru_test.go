package LRU

import (
	"fmt"
)

var lru = Constructor[int, int](5)

func init() {
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	lru.Put(5, 5)
	fmt.Println(lru.head.next)
	lru.Get(1)
	fmt.Println(lru.head.next)
}
