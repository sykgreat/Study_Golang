package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x: // 可写
			x = y
			y = x + y
		case <-quit: // 可读
			fmt.Println("Quitting...")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// sub goroutine
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	// main goroutine
	fibonacii(c, quit)
}
