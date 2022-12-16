package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 3) // 带有缓冲的channel
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("子 goroutine 结束。。。")

		for i := 1; i <= 4; i++ {
			c <- "fyy"
			fmt.Println("子 goroutine 正在运行。。。", "发送元素 i = ", i, " len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("main goroutine 结束")
}
