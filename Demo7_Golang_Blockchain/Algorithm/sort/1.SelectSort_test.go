package sort

import (
	"testing"
)

func Benchmark_SelectSort(b *testing.B) {
	arr := []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	selectSort(arr)
	b.Log(arr)
}

func Benchmark_SelectSort_String(b *testing.B) {
	arr := []string{"g", "a", "b", "c", "d", "e", "f"}
	selectSort(arr)
	b.Log(arr)
}

func selectSort[T int | string](arr []T) {
	for i := 0; i < len(arr)-1; i++ {
		index := i                          // index用于记录最大值的下标
		for j := i + 1; j < len(arr); j++ { // 从i+1开始，因为i是最大值的下标
			if arr[index] > arr[j] {
				index = j // 如果有比当前最大值更大的元素，就更新最大值的下标
			}
		}
		if index != i {
			arr[i], arr[index] = arr[index], arr[i] // 交换
		}
	}
}
