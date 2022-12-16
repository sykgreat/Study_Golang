package main

import "fmt"

func main() {
	// 定义一个channel
	c := make(chan string) // 带无缓冲的channel

	go func() {
		defer fmt.Println("goroutine 结束。。。")
		fmt.Println("goroutine 正在运行。。。")

		c <- "fyy" // 将 fyy 发送给c
	}()

	str := <-c // 从c中接受数据 并赋值给str

	fmt.Println("str = ", str)
	fmt.Println("main goroutine 结束。。。")
}
