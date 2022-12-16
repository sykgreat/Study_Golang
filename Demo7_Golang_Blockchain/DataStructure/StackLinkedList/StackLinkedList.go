package StackLinkedList

import (
	"errors"
	"fmt"
)

// MyStackLinkedList StackLinkedList接口
type MyStackLinkedList interface {
	Size() int               // 获取栈中元素个数
	IsEmpty() bool           // 判断栈是否为空
	Peek() interface{}       // 获取栈顶元素
	Push(value interface{})  // 入栈
	Pop() interface{}        // 出栈
	GetValue() []interface{} // 获取栈中所有元素
	Print()                  // 打印栈中所有元素
}

// SNode 链表栈节点
type SNode struct {
	value interface{}
	next  *SNode
}

func NewSNode(value interface{}, node *SNode) *SNode {
	return &SNode{
		value,
		node,
	}
}

// StackLinkedList 链表栈
type StackLinkedList struct {
	head *SNode
	size int
}

// NewStackLinkedList 构造函数
func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{
		head: nil,
		size: 0,
	}
}

// Size 获取栈中元素个数
func (s *StackLinkedList) Size() int {
	return s.size
}

// IsEmpty 判断栈是否为空
func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

// Peek 获取栈顶元素
func (s *StackLinkedList) Peek() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack Empty Exception"))
		return 0
	}
	return s.head.value
}

// Push 入栈
func (s *StackLinkedList) Push(value interface{}) {
	s.head = NewSNode(value, s.head)
	s.size++
}

// Pop 出栈
func (s *StackLinkedList) Pop() interface{} {
	if s.IsEmpty() {
		panic(errors.New("stack Empty Exception"))
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value
}

// GetValue 获取栈中所有元素
func (s *StackLinkedList) GetValue() (values []interface{}) {
	temp := s.head
	for temp != nil {
		values = append(values, temp.value)
		temp = temp.next
	}
	return values
}

// Print 打印栈中所有元素
func (s *StackLinkedList) Print() {
	temp := s.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}
