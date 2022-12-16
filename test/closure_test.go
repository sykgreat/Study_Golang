package test

import (
	"fmt"
	"testing"
)

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, v := range values {
		fmt.Printf("foo3 val = %d\n", v)
	}
}

func foo4() {
	values := []int{1, 2, 3, 5}
	for _, v := range values {
		go show(v)
	}
}

func show(v interface{}) {
	fmt.Printf("foo4 val = %d\n", v)
}

func foo5() {
	values := []int{1, 2, 3, 5}
	for _, v := range values {
		go func() {
			fmt.Printf("foo5 val = %d\n", v)
		}()
	}
}

var foo6Chan = make(chan int, 10)

func foo6() {
	for v := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", v)
		}()
	}
}

func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, v := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+v)
		})
	}
	return fs
}

func TestClosureQ1(t *testing.T) {
	// 第一组
	fmt.Println("第一组")
	x := 133
	f1 := foo1(&x)
	f2 := foo2(x)
	f1()
	f2()
	f1()
	f2()
	fmt.Println()

	// 第二组
	fmt.Println("第二组")
	x = 233
	f1()
	f2()
	f1()
	f2()
	fmt.Println()

	// 第三组a
	fmt.Println("第三组")
	x = 233
	foo1(&x)()
	foo2(x)()
	foo1(&x)()
	foo2(x)()
	fmt.Println()
}
