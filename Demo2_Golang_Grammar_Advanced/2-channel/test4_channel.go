package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // 关闭channel
	}()

	// 可以使用 range 用来迭代不断操作的channel
	for data := range c {
		fmt.Println(data)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Main Finished!!!")
}
