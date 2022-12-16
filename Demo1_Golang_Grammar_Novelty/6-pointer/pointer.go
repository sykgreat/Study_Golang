package main

import "fmt"

//func swap(a int, b int) {
//	temp := a
//	a = b
//	b = temp
//}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a = ", a, "b = ", b)
	fmt.Println()

	p := &a
	fmt.Println("p = ", p)
	fmt.Println("a = ", &a)
	fmt.Println()

	var pp **int // 二级指针
	pp = &p
	fmt.Println("pp = ", pp)
	fmt.Println("p  = ", &p)
}
