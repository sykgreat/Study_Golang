package linkedlist

import "Study_Golang/utils"

type LinkedNode[T any] struct {
	val  T
	next *LinkedNode[T]
}

func NewLinkedNode[T any](val T) *LinkedNode[T] {
	return &LinkedNode[T]{
		val:  val,
		next: nil,
	}
}

func (ln *LinkedNode[T]) addNode(val T) {
	for ln.next != nil {
		ln = ln.next
	}
	ln.next = NewLinkedNode[T](val)
}

func (ln *LinkedNode[T]) print() {
	str := ""
	for ln != nil {
		e, err := utils.ToStringE(ln.val)
		if err != nil {
			panic(err)
		}
		ln = ln.next
		if ln != nil {
			str += e + " --> "
		} else {
			str += e
		}
	}
	println(str)
}

func reverseList[T any](head *LinkedNode[T]) *LinkedNode[T] {
	var pre *LinkedNode[T]
	cur := head
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
