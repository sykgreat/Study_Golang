package main

import (
	"Study_Golang/Demo1_Golang_Grammar_Novelty/7-defer/demo"
	"fmt"
)

func main() {
	// 写入defer关键字
	defer fmt.Println("xhend 1")
	defer fmt.Println("xhend 2")

	fmt.Println("main:: sykxhfyy ::1")
	fmt.Println("main:: sykxhfyy ::2")
	fmt.Println()

	defer fmt.Println()
	defer demo.S()
	defer demo.Y()
	defer demo.K()

	demo.DeferAndReturnFunc()
	fmt.Println()
}
