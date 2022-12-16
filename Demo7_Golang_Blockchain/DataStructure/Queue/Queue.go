package Queue

import "errors"

// MyQueue MyQueue接口
type MyQueue interface {
	Size() int                // 返回队列的大小
	EnQueue(data interface{}) // 入队
	DeQueue() interface{}     // 出队
	IsEmpty() bool            // 判断队列是否为空
	Front() interface{}       // 返回队列的第一个元素
	End() interface{}         // 返回队列的最后一个元素
	Clear()                   // 清空队列
}

// Queue 队列
type Queue struct {
	dataSource []interface{} // 存储数据的切片
	theSize    int           // 队列的大小
}

// NewQueue 构造函数
func NewQueue() *Queue {
	return &Queue{
		dataSource: make([]interface{}, 0),
		theSize:    0,
	}
}

// Size 返回队列的大小
func (queue *Queue) Size() int {
	return queue.theSize
}

// EnQueue 入队
func (queue *Queue) EnQueue(data interface{}) {
	queue.dataSource = append(queue.dataSource, data)
	queue.theSize++
}

// DeQueue 出队
func (queue *Queue) DeQueue() interface{} {
	if queue.IsEmpty() {
		panic(errors.New("queue is empty"))
	}
	data := queue.dataSource[0]
	queue.dataSource = queue.dataSource[1:]
	queue.theSize--
	return data
}

// IsEmpty 判断队列是否为空
func (queue *Queue) IsEmpty() bool {
	return queue.theSize == 0
}

// Front 返回队列的第一个元素
func (queue *Queue) Front() interface{} {
	if queue.IsEmpty() {
		panic(errors.New("queue is empty"))
	}
	return queue.dataSource[0]
}

// End 返回队列的最后一个元素
func (queue *Queue) End() interface{} {
	if queue.IsEmpty() {
		panic(errors.New("queue is empty"))
	}
	return queue.dataSource[queue.theSize-1]
}

// Clear 清空队列
func (queue *Queue) Clear() {
	queue.dataSource = make([]interface{}, 0)
	queue.theSize = 0
}
