package QueueLinkedList

import "testing"

var queue = NewQueueLinkedList()

func init() {
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
}

func TestQueueLinkedList_Size(t *testing.T) {
	t.Log(queue.Size())
}

func TestQueueLinkedList_IsEmpty(t *testing.T) {
	t.Log(queue.IsEmpty())
}

func TestQueueLinkedList_EnQueue(t *testing.T) {
	queue.EnQueue(5)
	t.Log(queue.GetValue())
}

func TestQueueLinkedList_DeQueue(t *testing.T) {
	queue.DeQueue()
	t.Log(queue.GetValue())
}

func TestQueueLinkedList_Peek(t *testing.T) {
	t.Log(queue.Peek())
}

func TestQueueLinkedList_GetValue(t *testing.T) {
	t.Log(queue.GetValue())
}

func TestQueueLinkedList_Print(t *testing.T) {
	queue.Print()
}
