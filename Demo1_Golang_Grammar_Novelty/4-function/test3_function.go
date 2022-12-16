package main

import "fmt"

// 返回一个返回值
func fool(a string, b string) string {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return a + "like" + b
}

// 返回多个返回值， 匿名的
func fool2(a string, b string) (string, string) {
	return fool(a, b), fool(b, a)
}

// 返回多个返回值， 有形参名称的
func fool3(a string, b string) (r1 string, r2 string) { // ====> func fool3(a string, b string) (r1, r2 string) {

	// r1, r2 属于fool3的形参 值为默认值
	// r1, r2 作用空间 是fool3整个函数的{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = fool(a, b)
	r2 = fool(b, a)
	return
}

func main() {
	c := fool("syk", "ffy")
	fmt.Println("c = ", c)
	fmt.Println()

	ret1, ret2 := fool2("syk", "ffy")
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)
	fmt.Println()

	ret3, ret4 := fool3("sky", "ffy")
	fmt.Println("ret3 = ", ret3, "ret4 = ", ret4)
	fmt.Println()
}
