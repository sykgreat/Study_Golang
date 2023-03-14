package test

import (
	"sync"
	"testing"
)

func Test_ConcurrentAppendSliceNotForceIndex(t *testing.T) {
	sl := make([]int, 0)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sl = append(sl, num)
		}(k)
	}
	wg.Wait()
	t.Log(len(sl))
	t.Log(cap(sl))
}

func Test_ConcurrentAppendSliceForceIndex(t *testing.T) {
	sl := make([]int, 100)
	wg := sync.WaitGroup{}
	for index := 0; index < 100; index++ {
		k := index
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sl[num] = num
		}(k)
	}
	wg.Wait()
	t.Log(len(sl))
	t.Log(cap(sl))
}

// (512 + (512 + 256 * 3) / 4) * 8 = 6656 查表（/runtime/sizeclasses.go） 6784 / 8 = 848
// (848 + (848 + 256 * 3) / 4) * 8 = 10016 查表（/runtime/sizeclasses.go） 10240 / 8 = 1280
func Test_Expansion_one_int64(t *testing.T) {
	arr := make([]int64, 0)
	for i := 0; i < 1025; i++ {
		arr = append(arr, int64(i))
		if i%128 == 0 {
			t.Logf("len: %d, cap: %d", len(arr), cap(arr))
		}
	}
}

// (512 + (512 + 256 * 3) / 4) * 4 = 3328 查表（/runtime/sizeclasses.go） 3456 / 4 = 864
// (848 + (848 + 256 * 3) / 4) * 4 = 5008 查表（/runtime/sizeclasses.go） 5376 / 4 = 1344
func Test_Expansion_one_int32(t *testing.T) {
	arr := make([]int32, 0)
	for i := 0; i < 1025; i++ {
		arr = append(arr, int32(i))
		if i%128 == 0 {
			t.Logf("len: %d, cap: %d", len(arr), cap(arr))
		}
	}
}

// 1025 * 8 = 8200 查表（/runtime/sizeclasses.go） 9472 / 8 = 1184
func Test_Expansion_batch_append(t *testing.T) {
	arr1 := make([]int64, 0)
	arr2 := make([]int64, 0)

	for i := 0; i < 1025; i++ {
		arr1 = append(arr1, int64(i))
	}
	arr2 = append(arr2, arr1...)
	t.Logf("len: %d, cap: %d", len(arr1), cap(arr1))
	t.Logf("len: %d, cap: %d", len(arr2), cap(arr2))

}

func Test_array(t *testing.T) {
	arr := [3]int{1, 2, 3}
	for k, v := range arr {
		if k == 0 {
			arr[0], arr[1] = 100, 200
			t.Log(arr)
		}
		arr[k] = v + 100
	}
	t.Log(arr)
}

func Test_slice(t *testing.T) {
	arr := []int{1, 2, 3}
	for k, v := range arr {
		if k == 0 {
			arr[0], arr[1] = 100, 200
			t.Log(arr)
		}
		arr[k] = v + 100
	}
	t.Log(arr)
}
