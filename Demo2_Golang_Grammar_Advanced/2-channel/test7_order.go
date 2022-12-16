package main

import (
	"fmt"
)

type Order struct {
	id int
}

func producerOrder(out chan<- Order) {
	for i := 0; i < 10; i++ {
		data := Order{i}
		fmt.Println("生成了一个订单，订单ID为：", data.id)
		out <- data
	}
	close(out)
}

func consumerOrder(in <-chan Order, chQ chan bool) {
	for data := range in {
		fmt.Println("消费了一个订单，订单ID为：", data.id)
	}
	chQ <- true
}

func main() {
	ch := make(chan Order, 1)
	chQ := make(chan bool)
	go producerOrder(ch)
	go consumerOrder(ch, chQ)

	select {
	case _ = <-chQ:
		fmt.Println("消费完成!!!")
	}
}
