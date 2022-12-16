package main

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		data := i * i
		fmt.Println("生产者生产了一个数据：", data)
		out <- data
	}
	close(out)
}

func consumer(in <-chan int) {
	for data := range in {
		fmt.Println("消费者消费了一个数据：", data)
	}
}

func main() {
	// 创建一个无缓冲的channel
	//ch := make(chan int)
	// 创建一个有缓冲的channel
	ch := make(chan int, 5)
	go producer(ch)
	consumer(ch)
}
