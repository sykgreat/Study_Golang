package main

import (
	"fmt"
)

// interface{} 在编译期间无法确定其具体的参数类型，所以需要在运行时动态分配内存（内存分配到堆中）
func main() {
	a := 666
	fmt.Println(a) // println接收的类型为interface{}，所以a会被转换为interface{}类型，这个过程会发生逃逸
}
