package main

import (
	"Study_Golang/Demo1_Golang_Grammar_Novelty/5-init/lib1"
	"Study_Golang/Demo1_Golang_Grammar_Novelty/5-init/lib2"
	//_ "Study_Golang/Demo1_Golang_Grammar_Novelty/5-init/lib2" // 1 匿名导包 可以进行 init（）
	//. "Study_Golang/Demo1_Golang_Grammar_Novelty/5-init/lib2" // 2 将 lib2 中的 func 全部导入
	//myliib2 "Study_Golang/Demo1_Golang_Grammar_Novelty/5-init/lib2" // 3 为 lib2 取别名
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
	// 2 Lib2Test()
	// 3 myliib2.Lib2Test()
}
