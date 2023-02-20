package DeLinkList

import "fmt"

type MyDeLinkList[T comparable] interface {
	Size() int                 // 获取链表中元素个数
	IsEmpty() bool             // 判断链表是否为空
	GetValue() []any           // 获取链表中所有元素
	InsertFirst(value T)       // 在链表头插入元素
	InsertLast(value T)        // 在链表尾插入元素
	Insert(index int, value T) // 在指定位置插入元素
	RemoveFirst() any          // 删除链表头元素
	RemoveLast() any           // 删除链表尾元素
	Remove(index int) any      // 删除指定位置的元素
	Contains(value T) bool     // 判断链表中是否包含某个元素
	Find(value T) int          // 查找链表中某个元素的位置
	Reverse()                  // 反转链表
	Print()                    // 打印链表中所有元素
	Clear()                    // 清空链表
}

// DNode 双向链表节点
type DNode[T comparable] struct {
	value T
	prev  *DNode[T]
	next  *DNode[T]
}

// NewDNode 构造函数
func NewDNode[T comparable](value T, prev, next *DNode[T]) *DNode[T] {
	return &DNode[T]{
		value,
		prev,
		next,
	}
}

// DeLinkList 双向链表
type DeLinkList[T comparable] struct {
	head *DNode[T]
	rear *DNode[T]
	size int
}

// NewDeLinkList 构造函数
func NewDeLinkList[T comparable]() *DeLinkList[T] {
	return &DeLinkList[T]{
		head: nil,
		rear: nil,
		size: 0,
	}
}

// Size 获取链表中元素个数
func (dll *DeLinkList[T]) Size() int {
	return dll.size
}

// IsEmpty 判断链表是否为空
func (dll *DeLinkList[T]) IsEmpty() bool {
	return dll.size == 0
}

// GetValue 获取链表中所有元素
func (dll *DeLinkList[T]) GetValue() []T {
	var values []T
	cur := dll.head
	for cur != nil {
		values = append(values, cur.value)
		cur = cur.next
	}
	return values
}

// InsertFirst 在链表头插入元素
func (dll *DeLinkList[T]) InsertFirst(value T) {
	if dll.IsEmpty() {
		dll.head = NewDNode[T](value, nil, nil)
		dll.rear = dll.head
	} else {
		node := NewDNode[T](value, nil, dll.head)
		dll.head.prev = node
		dll.head = node
	}
	dll.size++
}

// InsertLast 在链表尾插入元素
func (dll *DeLinkList[T]) InsertLast(value T) {
	if dll.IsEmpty() {
		dll.head = NewDNode[T](value, nil, nil)
		dll.rear = dll.head
	} else {
		node := NewDNode[T](value, dll.rear, nil)
		dll.rear.next = node
		dll.rear = node
	}
	dll.size++
}

// Insert 在指定位置插入元素
func (dll *DeLinkList[T]) Insert(index int, value T) {
	if index < 0 || index > dll.size+1 {
		panic("index out of range")
	}
	if index == 0 {
		dll.InsertFirst(value)
	} else if index == dll.size+1 {
		dll.InsertLast(value)
	} else {
		mid := dll.size / 2
		if index <= mid {
			cur := dll.head
			for i := 0; i < index; i++ {
				cur = cur.next
			}
			node := NewDNode[T](value, cur.prev, cur)
			cur.prev.next = node
			cur.prev = node
			dll.size++
		} else {
			cur := dll.rear
			for i := dll.size - 1; i > index; i-- {
				cur = cur.prev
			}
			node := NewDNode[T](value, cur.prev, cur)
			cur.prev.next = node
			cur.prev = node
			dll.size++
		}
	}
}

// RemoveFirst 删除链表头元素
func (dll *DeLinkList[T]) RemoveFirst() T {
	if dll.IsEmpty() {
		panic("linklist is empty")
	}
	value := dll.head.value
	if dll.size == 1 {
		dll.head = nil
		dll.rear = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
	dll.size--
	return value
}

// RemoveLast 删除链表尾元素
func (dll *DeLinkList[T]) RemoveLast() T {
	if dll.IsEmpty() {
		panic("linklist is empty")
	}
	value := dll.rear.value
	if dll.size == 1 {
		dll.head = nil
		dll.rear = nil
	} else {
		dll.rear = dll.rear.prev
		dll.rear.next = nil
	}
	dll.size--
	return value
}

// Remove 删除指定位置的元素
func (dll *DeLinkList[T]) Remove(index int) T {
	if dll.IsEmpty() {
		panic("linklist is empty")
	}
	if index < 0 || index >= dll.size {
		panic("index out of range")
	}
	if index == 0 {
		return dll.RemoveFirst()
	} else if index == dll.size-1 {
		return dll.RemoveLast()
	} else {
		mid := dll.size / 2
		if index <= mid {
			cur := dll.head
			for i := 0; i < index; i++ {
				cur = cur.next
			}
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			dll.size--
			return cur.value
		} else {
			cur := dll.rear
			for i := dll.size - 1; i > index; i-- {
				cur = cur.prev
			}
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			dll.size--
			return cur.value
		}
	}
}

// Contains 判断链表中是否包含某个元素
func (dll *DeLinkList[T]) Contains(value T) bool {
	cur := dll.head
	for cur != nil {

		if cur.value == value {
			return true
		}
		cur = cur.next
	}
	return false
}

// Find 查找链表中某个元素的位置
func (dll *DeLinkList[T]) Find(value T) int {
	cur := dll.head
	for i := 0; i < dll.size; i++ {
		if cur.value == value {
			return i
		}
		cur = cur.next
	}
	return -1
}

// Clear 清空链表
func (dll *DeLinkList[T]) Clear() {
	dll.head = nil
	dll.rear = nil
	dll.size = 0
}

// Reverse 反转链表
func (dll *DeLinkList[T]) Reverse() {
	if dll.IsEmpty() {
		panic("linklist is empty")
	}
	cur := dll.head
	for cur != nil {
		cur.prev, cur.next = cur.next, cur.prev
		cur = cur.prev
	}
	dll.head, dll.rear = dll.rear, dll.head
}

// Print 打印链表
func (dll *DeLinkList[T]) Print() {
	cur := dll.head
	for cur != nil {
		fmt.Printf("%v ", cur.value)
		cur = cur.next
	}
	fmt.Println()
}
