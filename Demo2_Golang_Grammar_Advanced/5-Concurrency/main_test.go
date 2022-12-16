package Concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	testMap = make(map[int]int, 10)
	lock    sync.Mutex
)

func Test_concurrency(t *testing.T) {
	start := time.Now()
	for i := 1; i <= 20; i++ {
		go func(num int) {
			res := 1
			for i := 1; i <= num; i++ {
				res *= i
			}
			testMap[num] = res
		}(i)
	}

	time.Sleep(time.Second * 5)
	for key, val := range testMap {
		fmt.Printf("数字%v对应的阶乘是: %v\n", key, val)
	}
	end := time.Since(start)
	fmt.Println(end)
}

func Test_concurrency_lock(t *testing.T) {
	start := time.Now()
	for i := 1; i <= 20; i++ {
		go func(num int) {
			res := 1
			for i := 1; i <= num; i++ {
				res *= i
			}
			lock.Lock()
			defer lock.Unlock()
			testMap[num] = res
		}(i)
	}

	time.Sleep(time.Second * 5)
	for key, val := range testMap {
		fmt.Printf("数字%v对应的阶乘是: %v\n", key, val)
	}
	end := time.Since(start)
	fmt.Println(end)
}
