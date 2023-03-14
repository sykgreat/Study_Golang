package linkedlist

import "testing"

var linkedList = NewLinkedNode[int](1)

func init() {
	linkedList.addNode(2)
	linkedList.addNode(3)
	linkedList.addNode(4)
	linkedList.print()
}

func TestReverseList(t *testing.T) {
	head := reverseList(linkedList)
	head.print()
}
