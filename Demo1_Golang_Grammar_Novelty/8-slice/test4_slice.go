package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加元素1，numbers len = 4, cap = 5, slice = [0 0 0 1]
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加元素2，numbers len = 5, cap = 5, slice = [0 0 0 1 1]
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向一个容量cap已经满的slice追加元素， numbers len = 6, cap = 10, slice = [0 0 0 1 2 3]
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	fmt.Println()
}
