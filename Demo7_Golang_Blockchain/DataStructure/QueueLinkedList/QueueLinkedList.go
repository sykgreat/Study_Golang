package QueueLinkedList

import "fmt"

// MyQueueLinkedList QueueLinkedList接口
type MyQueueLinkedList interface {
	Size() int               // 队列大小
	IsEmpty() bool           // 队列是否为空
	EnQueue(value int)       // 入队
	DeQueue() int            // 出队
	Peek() int               // 获取队首元素
	GetValue() []interface{} // 获取队列中所有元素
	Print()                  // 打印队列中所有元素
}

// QNode 链表队列节点
type QNode struct {
	data int
	next *QNode
}

// NewQNode 构造函数
func NewQNode(value int) *QNode {
	return &QNode{
		value,
		nil,
	}
}

// QueueLinkedList 链表队列
type QueueLinkedList struct {
	head  *QNode
	tail  *QNode
	count int
}

// NewQueueLinkedList 构造函数
func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{
		nil,
		nil,
		0,
	}
}

// Size 获取队列大小
func (q QueueLinkedList) Size() int {
	return q.count
}

// IsEmpty 判断队列是否为空
func (q QueueLinkedList) IsEmpty() bool {
	return q.count == 0
}

// EnQueue 入队
func (q *QueueLinkedList) EnQueue(value int) {
	node := NewQNode(value)
	if q.head == nil {
		q.head = node
	} else {
		q.tail.next = node
	}
	q.count++
	q.tail = node
}

// DeQueue 出队
func (q *QueueLinkedList) DeQueue() int {
	if q.head == nil {
		fmt.Println("Empty Queue")
		return -1
	}
	temp := q.head
	q.head = q.head.next
	q.count--
	if q.head == nil {
		q.tail = nil
	}
	return temp.data
}

// Peek 获取队首元素
func (q *QueueLinkedList) Peek() int {
	if q.head == nil {
		fmt.Println("Empty Queue")
		return -1
	}
	return q.head.data
}

// GetValue 获取队列中所有元素
func (q *QueueLinkedList) GetValue() (result []interface{}) {
	var temp *QNode = q.head
	for temp != nil {
		result = append(result, temp.data)
		temp = temp.next
	}
	return result
}

// Print 打印队列中所有元素
func (q *QueueLinkedList) Print() {
	var temp *QNode = q.head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}
