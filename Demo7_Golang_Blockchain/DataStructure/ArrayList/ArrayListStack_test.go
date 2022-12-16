package ArrayList

import "testing"

var stack StackArray = NewArrayListStack(10)

func init() {
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
}

func TestStack_Clear(t *testing.T) {
	stack.Clear()
	t.Log(stack.Size())
}

func TestStack_Size(t *testing.T) {
	t.Log(stack.Size())
}

func TestStack_Pop(t *testing.T) {
	t.Log(stack.Pop())
}

func TestStack_Push(t *testing.T) {
	stack.Push(5)
	t.Log(stack.Size())
}

func TestStack_IsFull(t *testing.T) {
	t.Log(stack.IsFull())
}

func TestStack_IsEmpty(t *testing.T) {
	t.Log(stack.IsEmpty())
}
