package Interview

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func Test1(t *testing.T) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(3)

	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(1)

	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(2)

	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(3)

	wg.Wait()
}

func Test2(t *testing.T) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	for i := 1; i <= 258; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func Test3(t *testing.T) {
	go bzd()()
}

func bzd() (wsm func()) {
	wsm = func() {
		fmt.Println("wsm")
	}
	return func() {
		fmt.Println("nss")
		wsm()
	}
}
