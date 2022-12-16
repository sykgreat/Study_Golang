package Queue

import "testing"

var queue = NewQueue()

func init() {
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
}

func TestQueue_Size(t *testing.T) {
	t.Log(queue.Size())
}

func TestQueue_Front(t *testing.T) {
	t.Log(queue.Front())
}

func TestQueue_End(t *testing.T) {
	t.Log(queue.End())
}

func TestQueue_EnQueue(t *testing.T) {
	queue.EnQueue(5)
	t.Log(queue.dataSource)
}

func TestQueue_DeQueue(t *testing.T) {
	t.Log(queue.DeQueue())
	t.Log(queue.dataSource)
}

func TestQueue_IsEmpty(t *testing.T) {
	t.Log(queue.IsEmpty())
}

func TestQueue_Clear(t *testing.T) {
	queue.Clear()
	t.Log(queue.dataSource)
}
