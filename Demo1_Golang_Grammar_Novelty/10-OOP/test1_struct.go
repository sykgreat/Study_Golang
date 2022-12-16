package main

import "fmt"

// 声明一种新的数据类型 myInt ，是int的一个别名
type myInt int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) { // 传递副本
	book.auth = "syk"
}

func changeBook2(book *Book) { // 指针传递
	book.auth = "syk"
}

func main() {
	var a myInt = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)
	fmt.Println()

	var book1 Book
	book1.title = "Golang"
	book1.auth = "hpq"
	fmt.Println("book1 = ", book1)
	fmt.Println()

	changeBook1(book1)
	fmt.Println("book1 = ", book1)
	fmt.Println()

	changeBook2(&book1)
	fmt.Println("book1 = ", book1)
	fmt.Println()
}
