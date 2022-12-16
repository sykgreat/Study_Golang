package main

import "fmt"

func main() {
	// 声明slice1是一个切片，并初始化，默认值是1，2，3。长度len是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("slice1 len = %d, slice1 = %v\n", len(slice1), slice1)
	fmt.Println()

	// 声明slice2是一个切片，但是并没有给slice分配空间
	var slice2 []int
	// slice2 = make([]int, 3) // 给slice2开辟空间
	fmt.Printf("slice2 len = %d, slice2 = %v\n", len(slice2), slice2)
	fmt.Println()

	// 声明slice3是一个切片，同时给slice3分配空间
	var slice3 []int = make([]int, 3) // ===》 slice3 := make([]int, 3) 最常用
	fmt.Printf("slice3 len = %d, slice3 = %v\n", len(slice3), slice3)
	fmt.Println()

	isNil(slice2)
}

// 判断切片是否为空
func isNil(slice []int) {
	if slice == nil {
		fmt.Println("slice 是一个空切片")
	} else {
		fmt.Println("slice 是有空间的")
	}
}
