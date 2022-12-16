package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型

type Book struct{}

func (b *Book) ReadBook() {
	fmt.Println("read a book...")
}

func (b *Book) WriteBook() {
	fmt.Println("write a book...")
}

func main() {
	// b: pair<type: *Book, value: book{}地址>
	b := &Book{}

	// r: pair<type: , value: >
	var r Reader
	// r: pair<type: *Book, value: book{}地址>
	r = b
	r.ReadBook()

	// w: pair<type: , value: >
	var w Writer
	// w: pair<type: *Book, value: book{}地址>
	w = b
	w.WriteBook()

	w = r.(Writer) // 此处的断言为什么会成功？ 因为 w r具体的type是一致的
}
