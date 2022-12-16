package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{} 该如何区分 此时引用的底层数据类型到底是声明呢？
	// 给interface{}提供了 “断言” 机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("value is string type, value = ", value)
	}
	fmt.Println()
}

type myBook struct {
	Auth string
}

func main() {
	book := myBook{"Golang"}
	myFunc(book)
	myFunc(100)
	myFunc(3.14)
	myFunc("sykxhfyy")
}
