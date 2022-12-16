package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}

	// [0:2)
	s1 := slice1[0:2] //[1 2]
	fmt.Println(s1)

	// [0:3)
	s2 := slice1[:3] // [1 2 3]
	fmt.Println(s2)

	// [3:5]
	s3 := slice1[3:] // [4 5 6]
	fmt.Println(s3)
	fmt.Println()

	//copy 可以将底层数组的slice一起进行拷贝
	slice2 := make([]int, 3)

	copy(slice2, slice1) // 将slice1中的值 依次拷贝到slice2
	fmt.Println(slice2)
}
