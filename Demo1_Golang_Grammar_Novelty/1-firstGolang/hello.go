package main // 程序的包名字

import (
	"fmt"
	"time"
)

// main 主函数
func main() { // 函数的 { 一定和函数名在同一行，否则编译错误
	// golang中的表达式，加“；”，和不加“；”都可以，建议不加
	fmt.Println("hello fyy")

	time.Sleep(10 * time.Second)
}
