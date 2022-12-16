package Algorithm

import (
	"testing"
)

func Benchmark_BubbleSort(b *testing.B) {
	arr := []int{10, 1, 8, 3, 9, 4, 2, 5, 7, 6}
	bubbleSort(arr)
	b.Log(arr)
}

func Benchmark_BubbleSort_String(b *testing.B) {
	arr := []string{"g", "a", "c", "e", "d", "f", "b"}
	bubbleSort(arr)
	b.Log(arr)
}

func bubbleSort[T int | string](arr []T) {
	for i := 0; i < len(arr)-1; i++ { // 从0开始，因为最后一个元素不需要比较
		isNeedChange := false
		for j := 0; j < len(arr)-1-i; j++ { // 从0开始，因为最后一个元素不需要比较
			if arr[j] > arr[j+1] { // 如果前一个元素比后一个元素大，就交换
				arr[j], arr[j+1] = arr[j+1], arr[j] // 交换
				isNeedChange = true
			}
		}
		if !isNeedChange {
			break
		}
	}
}
