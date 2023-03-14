package sort

import (
	"testing"
)

func Benchmark_HeapFindMin(b *testing.B) {
	arr := []int{10, 1, 8, 3, 9, 4, 2, 5, 7, 6}
	b.Log(heapFindMin(arr))
}

func heapFindMin[T int | string](arr []T) T {
	for i := len(arr)/2 - 1; i >= 0; i-- { // 假设当前节点为最小值
		leftChild := 2*i + 1                                 // 左孩子
		rightChild := 2*i + 2                                // 右孩子
		if leftChild < len(arr) && arr[leftChild] < arr[i] { // 如果左孩子小于当前节点
			arr[leftChild], arr[i] = arr[i], arr[leftChild]
		}
		if rightChild < len(arr) && arr[rightChild] < arr[i] { // 如果右孩子小于当前节点
			arr[rightChild], arr[i] = arr[i], arr[rightChild]
		}
	}
	return arr[0]
}

func Benchmark_HeapSort(b *testing.B) {
	arr := []int{10, 1, 8, 3, 9, 4, 2, 5, 7, 6}
	heapSort(arr)
	b.Log(arr)
}

func Benchmark_HeapSort_String(b *testing.B) {
	arr := []string{"g", "a", "b", "c", "d", "e", "f"}
	heapSort(arr)
	b.Log(arr)
}

func heapSort[T int | string](arr []T) {
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapIfy(arr, i)
	}
}

func heapIfy[T int | string](arr []T, unsortCapacity int) {
	for i := (unsortCapacity / 2) - 1; i >= 0; i-- {
		leftChild := 2*i + 1  // 左孩子
		rightChild := 2*i + 2 // 右孩子
		if leftChild < unsortCapacity && arr[i] < arr[leftChild] {
			arr[i], arr[leftChild] = arr[leftChild], arr[i]
		}
		if rightChild < unsortCapacity && arr[i] < arr[rightChild] {
			arr[i], arr[rightChild] = arr[rightChild], arr[i]
		}
	}
}
