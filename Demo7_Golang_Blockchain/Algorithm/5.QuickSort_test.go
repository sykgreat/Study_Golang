package Algorithm

import "testing"

func quickSort[T int | string](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	var left, right []T
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

func Benchmark_QuickSort(b *testing.B) {
	arr := []int{10, 1, 8, 3, 9, 4, 2, 5, 7, 6}
	sort := quickSort(arr)
	b.Log(sort)
}

func Benchmark_QuickSort_String(b *testing.B) {
	arr := []string{"g", "a", "b", "c", "d", "e", "f"}
	sort := quickSort(arr)
	b.Log(sort)
}
