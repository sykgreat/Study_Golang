package main

import "fmt"

func main() {
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}

	myArray3 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}
	fmt.Println()

	for i, v := range myArray2 {
		fmt.Println("index = ", i, "value = ", v)
	}
	fmt.Println()

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)
	fmt.Println()

	printArray(myArray3)
	fmt.Println()

	for i, v := range myArray2 {
		fmt.Println("index = ", i, "value = ", v)
	}
	fmt.Println()
}

func printArray(myArray [4]int) { // 值拷贝
	for i, v := range myArray {
		fmt.Println("index = ", i, "value = ", v)
	}
	myArray[0] = 10
}
