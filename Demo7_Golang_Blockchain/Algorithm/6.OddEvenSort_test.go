package Algorithm

import (
	"testing"
)

func Benchmark_OddEvenSort(b *testing.B) {
	arr := []int{10, 6, 1, 3, 2, 9, 4, 5, 8, 7}
	oddEvenSort(arr)
	b.Log(arr)
}

func Benchmark_OddEvenSort_String(b *testing.B) {
	arr := []string{"g", "a", "b", "c", "d", "e", "f"}
	oddEvenSort(arr)
	b.Log(arr)
}

func oddEvenSort[T int | string](arr []T) {
	isSorted := false
	for isSorted == false { // 当数据还需要排序的时候就是false
		isSorted = true
		for i := 1; i < len(arr)-1; i += 2 { // 奇数位
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}

		for i := 0; i < len(arr)-1; i += 2 { // 偶数位
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
}

func BenchmarkReOrderArrayV1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		reOrderArrayV1(arr) // run reOrderArrayV1(arr) b.N times
	}
}

func BenchmarkReOrderArrayV2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		reOrderArrayV2(arr, isEven) // run reOrderArrayV2(arr, isEven) b.N times
	}
}

func reOrderArrayV1(arr []int) []int {
	var oddArr, evenArr []int
	for _, value := range arr {
		if value%2 == 0 {
			evenArr = append(evenArr, value)
		} else {
			oddArr = append(oddArr, value)
		}
	}
	return append(oddArr, evenArr...)
}

// 根据指定闭包对数组切片排序
func reOrderArrayV2(arr []int, orderFunc func(int) bool) []int {
	// 小于等于1个元素无需处理
	if arr == nil || len(arr) <= 1 {
		return arr
	}
	// 设置两个指针，从头尾往中间移动
	i := 0
	j := len(arr) - 1
	// 头指针不能越过尾指针，否则退出
	// 以奇偶数排序为例，i 从左到右寻找偶数，j 从右到左寻找奇数
	// 该循环执行完毕后，i 左侧的都是奇数，j 右侧的都是偶数，也就完成了顺序调整
	for i < j {
		// 如果不符合条件，则头指针后移，否则中断
		// 以 orderFunc 为偶数判断函数为例，返回 false 表示是奇数
		// 题目要求奇数排在前面，因此，当 i 对应值是奇数时，往后移一位，然后继续下一个循环，直到 i==j 或者遇到第一个偶数中断
		for i < j && !orderFunc(arr[i]) {
			i++
		}
		// 如果符合条件，则尾指针前移，否则中断
		// 还是以 orderFunc 为偶数判断函数为例，返回 true 表示是偶数
		// 题目要求偶数排在后面，因此，当 j 对应值是偶数时，往前移一位，然后继续下一个循环，直到 j==i 或者遇到第一个奇数中断
		for i < j && orderFunc(arr[j]) {
			j--
		}
		// 如果 i < j，则交换对应值的位置
		// 以奇偶数为例，此时 arr[i] 是偶数，arr[j] 是奇数，则交换两个值，将奇数放到前面，偶数放到后面
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
		// 继续下一个循环，直到 i==j，此时 i 左侧都是奇数，j 右侧都是偶数，所有奇数都排到了偶数前面
	}
	return arr
}

// 排序条件：是否是偶数
func isEven(num int) bool {
	return num&1 == 0
}
