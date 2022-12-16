package ArrayList

import "errors"

// StackArray StackArray接口
type StackArray interface {
	Clear()                // 清空栈
	Size() int             // 返回栈的大小
	Pop() interface{}      // 弹出栈顶元素
	Push(data interface{}) // 压入元素到栈顶
	IsFull() bool          // 判断栈是否已满
	IsEmpty() bool         // 判断栈是否为空
}

// Stack 栈
type Stack struct {
	list    *ArrayList // 存储数据的切片
	capSize int        // 栈的容量
}

// NewArrayListStack 构造函数
func NewArrayListStack(capSize int) *Stack {
	return &Stack{
		list:    NewArrayList(10),
		capSize: capSize,
	}
}

// Clear 清空栈
func (stack *Stack) Clear() {
	stack.list.Clear()
}

// Size 返回栈的大小
func (stack *Stack) Size() int {
	return stack.list.Size()
}

// Pop 弹出栈顶元素
func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		panic(errors.New("stack is empty"))
	}
	data := stack.list.Get(stack.list.Size() - 1)
	stack.list.Delete(stack.list.Size() - 1)
	return data
}

// Push 压入元素到栈顶
func (stack *Stack) Push(data interface{}) {
	if stack.IsFull() {
		panic(errors.New("stack is full"))
	}
	stack.list.Append(data)
}

// IsFull 判断栈是否已满
func (stack *Stack) IsFull() bool {
	return stack.list.Size() >= stack.capSize
}

// IsEmpty 判断栈是否为空
func (stack *Stack) IsEmpty() bool {
	return stack.list.Size() == 0
}
