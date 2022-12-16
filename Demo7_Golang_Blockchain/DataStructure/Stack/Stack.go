package Stack

import "errors"

// MyStack Stack接口
type MyStack interface {
	Clear()                // 清空栈
	Size() int             // 返回栈的大小
	Pop() interface{}      // 弹出栈顶元素
	Push(data interface{}) // 压入元素到栈顶
	IsFull() bool          // 判断栈是否已满
	IsEmpty() bool         // 判断栈是否为空
}

// Stack 栈
type Stack struct {
	dataSource  []interface{} // 存储数据的切片
	capSize     int           // 栈的容量
	currentSize int           // 栈的当前大小
}

// NewStack 构造函数
func NewStack(capSize int) *Stack {
	return &Stack{
		dataSource:  make([]interface{}, 0, capSize),
		capSize:     capSize,
		currentSize: 0,
	}
}

// Clear 清空栈
func (stack *Stack) Clear() {
	stack.dataSource = make([]interface{}, stack.capSize)
	stack.currentSize = 0
}

// Size 返回栈的大小
func (stack *Stack) Size() int {
	return stack.currentSize
}

// Pop 弹出栈顶元素
func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	data := stack.dataSource[stack.currentSize-1]
	stack.dataSource = stack.dataSource[:stack.currentSize-1]
	stack.currentSize--
	return data
}

// Push 弹出栈顶元素
func (stack *Stack) Push(data interface{}) {
	if stack.IsFull() {
		panic(errors.New("stack is full"))
	}
	stack.dataSource = append(stack.dataSource, data)
	stack.currentSize++
}

// IsFull 判断栈是否已满
func (stack *Stack) IsFull() bool {
	return stack.currentSize >= stack.capSize
}

// IsEmpty 判断栈是否为空
func (stack *Stack) IsEmpty() bool {
	return stack.currentSize == 0
}
