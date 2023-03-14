package sort

import "testing"

func Benchmark_MergeSort(b *testing.B) {
	arr := []int{10, 6, 1, 3, 2, 9, 4, 5, 8, 7}
	b.Log(mergeSort(arr))
}

func Benchmark_MergeSort_stirng(b *testing.B) {
	arr := []string{"g", "a", "b", "c", "d", "e", "f"}
	b.Log(mergeSort(arr))
}

// 归并排序
func mergeSort[T int | string](nums []T) []T {
	if len(nums) <= 1 {
		return nums
	}

	p := len(nums) / 2 // 获取分区位置
	// 通过递归分区
	left := mergeSort(nums[0:p])
	right := mergeSort(nums[p:])

	return merge(left, right) // 排序后合并
}

// 排序合并
func merge[T int | string](left []T, right []T) []T {
	i, j := 0, 0
	m, n := len(left), len(right)

	var result []T // 用于存放结果集
	for i < m && j < n {
		// 对所有区间数据进行排序
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 如果左侧区间还没有遍历完，将剩余数据放到结果集
	for ; i < m; i++ {
		result = append(result, left[i])
	}

	// 如果右侧区间还没有遍历完，将剩余数据放到结果集
	for ; j < n; j++ {
		result = append(result, right[j])
	}

	return result
}
