package main

import "fmt"

const (
	// 可以在const()添加一个关键字iota，每行的iota都会累加1，第一行的iota的默认值为0

	BEIJING  = 10 * iota // iota = 0
	SHANGHAI             // iota = 10
	SHENZHEN             // iota = 20
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = 1, b = 2
	c, d                      // iota = 1, c = 2, d = 3
	e, f                      // iota = 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = 6, h = 9
	i, k                      // iota = 4 , i = 8, k = 12
)

func main() {
	// 常量（只读属性）
	const length int = 10
	fmt.Println("length = ", length)
	fmt.Println()
	//length = 100 // 常量不可以赋值

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
	fmt.Println()

	fmt.Println("a = ", a, " b = ", b)
	fmt.Println("c = ", c, " d = ", d)
	fmt.Println("e = ", e, " f = ", f)
	fmt.Println("g = ", g, " h = ", h)
	fmt.Println("i = ", i, " k = ", k)
}
