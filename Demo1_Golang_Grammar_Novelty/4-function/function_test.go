package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func Test_Closure(t *testing.T) {
	t.Log(Fibonacci(10)) // 1 1 2 3 5 8 13 21 34 55
}

func Fibonacci(x int) (res int) {
	if x == 1 || x == 2 {
		return 1
	}

	cf := closureFunc()
	for i := 3; i <= x; i++ {
		res = cf()
	}
	return
}

func closureFunc() func() int {
	x1 := 1
	x2 := 1
	x3 := 0
	return func() int {
		x3 = x1 + x2
		x1 = x2
		x2 = x3
		return x3
	}
}

func Test_ClosureCirculate(t *testing.T) {
	closureCirculate()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}

func closureCirculate() {
	// 闭包的坑
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
