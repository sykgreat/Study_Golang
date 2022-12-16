package main

import "fmt"

/*
	四种变量的声明方式
*/

// 声明全局变量，方法一、方法二、方法三是可以的，方法四不行

func main() {
	// 方法一：声明一个变量，默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)
	fmt.Println()

	// 方法二：声明一个变量，初始化一个值
	var b int = 10
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)
	fmt.Println()

	// 方法三：在初始化的时候，可以省去变量类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)
	fmt.Println()

	// 方法四：（常用方法）省去var关键字，直接自动匹配
	d := 1000
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)
	fmt.Println()

	// 声明多个变量
	var xx, yy int = 1, 2
	fmt.Println("xx = ", xx, "yy = ", yy)
	var aa, bb = 3, "ffy"
	fmt.Println("aa = ", aa, "bb = ", bb)
	var (
		cc = 4
		dd = true
	)
	fmt.Println("cc = ", cc, "dd = ", dd)
}
