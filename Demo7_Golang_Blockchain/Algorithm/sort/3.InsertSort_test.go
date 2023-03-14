package sort

import "testing"

func Benchmark_InsertSort(b *testing.B) {
	arr := []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	insertSort(arr)
	b.Log(arr)
}

func Benchmark_InsertSort_String(b *testing.B) {
	arr := []string{"g", "a", "b", "c", "d", "e", "f"}
	insertSort(arr)
	b.Log(arr)
}

func insertSort[T int | string](arr []T) {
	// 从第二个元素开始
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]  // 待插入的元素
		insertIndex := i - 1 // 待插入元素的前一个元素的下标
		// 如果待插入元素的前一个元素的下标大于等于0，且待插入元素的前一个元素的值大于待插入元素的值
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 将待插入元素的前一个元素的值赋值给待插入元素的后一个元素
			insertIndex--                         // 待插入元素的前一个元素的下标减1
		}
		arr[insertIndex+1] = insertVal // 将待插入元素的值赋值给待插入元素的前一个元素的后一个元素
	}
}
