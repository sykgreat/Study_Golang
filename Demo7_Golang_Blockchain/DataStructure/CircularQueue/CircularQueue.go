package CircularQueue

import "fmt"

// MyCircularQueue  MyCircularQueue接口
type MyCircularQueue interface {
	Size() int                  // 返回队列的大小
	IsEmpty() bool              // 判断队列是否为空
	IsFull() bool               // 判断队列是否已满
	EnQueue(v interface{}) bool // 入队
	DeQueue() interface{}       // 出队
	String() string             // 打印队列
}

// CircularQueue 环形队列
type CircularQueue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

// NewCircularQueue 构造函数
func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		data:     make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

// Size 返回队列的大小
func (q *CircularQueue) Size() int {
	return (q.tail - q.head + q.capacity) % q.capacity
}

// IsEmpty 判断队列是否为空
func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

// IsFull 判断队列是否已满
func (q *CircularQueue) IsFull() bool {
	return q.head == (q.tail+1)%q.capacity
}

// EnQueue 入队
func (q *CircularQueue) EnQueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}

	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

// DeQueue 出队
func (q *CircularQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) % q.capacity
	return v
}

// String 打印队列
func (q *CircularQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	var i = q.head
	for {
		result += fmt.Sprintf("<-%+v", q.data[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
