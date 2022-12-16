package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组，切片 slice
	fmt.Printf("myArray type is %T\n", myArray)

	printASlice(myArray)
	fmt.Println()

	for _, v := range myArray {
		fmt.Println("value = ", v)
	}
}

func printASlice(myArray []int) { // 引用传递
	for _, v := range myArray {
		fmt.Println("value = ", v)
	}
	myArray[0] = 10
}
