package CircularQueue

import "testing"

var queue = NewCircularQueue(10)

func init() {
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
}

func TestCircularQueue_Size(t *testing.T) {
	t.Log(queue.Size())
}

func TestCircularQueue_IsEmpty(t *testing.T) {
	t.Log(queue.IsEmpty())
}

func TestCircularQueue_IsFull(t *testing.T) {
	t.Log(queue.IsFull())
}

func TestCircularQueue_EnQueue(t *testing.T) {
	queue.EnQueue(5)
	t.Log(queue.data)
}

func TestCircularQueue_DeQueue(t *testing.T) {
	t.Log(queue.DeQueue())
	t.Log(queue.data)
}

func TestCircularQueue_String(t *testing.T) {
	t.Log(queue.String())
}
