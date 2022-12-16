package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// trace的编程过程
// 1 创建文件
// 2 启动
// 3 停止
func main() {
	// 1 创建一个trace文件
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// 2 启动trace
	err = trace.Start(file)
	if err != nil {
		panic(err)
	}

	// 正常要调试的业务
	fmt.Println("syk like fyy")

	// 3 停止trace
	trace.Stop()
}
