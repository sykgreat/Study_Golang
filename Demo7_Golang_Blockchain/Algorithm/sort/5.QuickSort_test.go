package sort

import (
	"golang.org/x/exp/constraints"
	"os"
	"strconv"
	"strings"
	"testing"
)

func quickSort[T constraints.Ordered](arr []T) []T {
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
	arr := make([]int, 0)
	dir, err := os.ReadDir("E:\\nginx-1.23.3\\html\\video\\KiminoNawa\\image")
	if err != nil {
		b.Log(err)
	}
	for _, v := range dir {
		if !v.IsDir() {
			str := strings.Split(v.Name(), ".")[0]
			parseInt, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				b.Log(err)
				return
			}
			arr = append(arr, int(parseInt))
		}
	}
	sort := quickSort(arr)
	b.Log(sort)
}

func Benchmark_QuickSort_String(b *testing.B) {
	arr := []string{"g", "a", "b", "c", "d", "e", "f"}
	sort := quickSort(arr)
	b.Log(sort)
}

func quickSortChannel[T constraints.Ordered](arr []T, c chan []T) {
	if len(arr) <= 1 {
		c <- arr
		return
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
	c1 := make(chan []T)
	c2 := make(chan []T)
	go quickSortChannel(left, c1)
	go quickSortChannel(right, c2)
	select {
	case left := <-c1:
		select {
		case right := <-c2:
			c <- append(append(left, pivot), right...)
		}
	}
}

func Benchmark_QuickSort_Channel(b *testing.B) {
	arr := make([]int, 0)
	dir, err := os.ReadDir("E:\\nginx-1.23.3\\html\\video\\KiminoNawa\\image")
	if err != nil {
		b.Log(err)
	}
	for _, v := range dir {
		if !v.IsDir() {
			str := strings.Split(v.Name(), ".")[0]
			parseInt, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				b.Log(err)
				return
			}
			arr = append(arr, int(parseInt))
		}
	}
	c := make(chan []int)
	go quickSortChannel(arr, c)
	select {
	case sort := <-c:
		b.Log(sort)
	}
}
