package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 用go创建承载一个形参为空 返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	go func(a, b string) string {
		fmt.Printf("%v love %v\n", a, b)
		return "yes" // 异步操作 拿不到返回值
	}("syk", "fyy")

	for {
		time.Sleep(1 * time.Second)
	}
}
