package StackLinkedList

import "testing"

var stack = NewStackLinkedList()

func init() {
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
}

func TestStackLinkedList_Size(t *testing.T) {
	t.Log(stack.Size())
}

func TestStackLinkedList_IsEmpty(t *testing.T) {
	t.Log(stack.IsEmpty())
}

func TestStackLinkedList_Peek(t *testing.T) {
	t.Log(stack.Peek())
	t.Log(stack.GetValue())
}

func TestStackLinkedList_Push(t *testing.T) {
	stack.Push(5)
	t.Log(stack.GetValue())
}

func TestStackLinkedList_Pop(t *testing.T) {
	t.Log(stack.Pop())
	t.Log(stack.GetValue())
}

func TestStackLinkedList_GetValue(t *testing.T) {
	t.Log(stack.GetValue())
}

func TestStackLinkedList_Print(t *testing.T) {
	stack.Print()
}
