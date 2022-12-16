package Algorithm

import (
	"testing"
)

func Benchmark_ShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{10, 6, 1, 3, 2, 9, 4, 5, 8, 7}
		shellSort(a)
		b.Log(a)
	}
}

func shellSort[T int | string](arr []T) {
	length := len(arr)
	if length == 1 {
		return
	} else {
		gap := length / 2
		for i := 0; i < gap; i++ {
			shellSortStep(arr, i, gap)
			gap /= 2
		}
	}

}

func shellSortStep[T int | string](arr []T, start int, gap int) {
	length := len(arr)

	// 插入排序的变种
	for i := start + gap; i < length; i += gap {
		insertVal := arr[i] // 备份待插入的数据
		j := i - gap        // 初始化待插入位置

		// 待插入数据，小于前面的数据
		for j >= 0 && insertVal < arr[j] {
			arr[j+gap] = arr[j] // 从前往后移动
			j -= gap
		}

		arr[j+gap] = insertVal
	}
}
