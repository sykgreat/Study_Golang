package test

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBzd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	var rchs []chan int
	for i := 0; i < 10; i++ {
		rchs = append(rchs, make(chan int))
	}

	// 创建SelectCase
	var cases = createRecvCases(rchs)

	// 消费者goroutine
	go func() {
		defer wg.Done()
		for {
			chosen, recv, ok := reflect.Select(cases)
			if ok {
				fmt.Printf("recv from channel [%d], val=%v\n", chosen, recv)
				continue
			}
			// one of the channels is closed, exit the goroutine
			fmt.Printf("channel [%d] closed, select goroutine exit\n", chosen)
			return
		}
	}()

	// 生产者goroutine
	go func() {
		defer wg.Done()
		var n int
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < 10; i++ {
			n = r.Intn(10)
			rchs[n] <- n
		}
		close(rchs[n])
	}()
	wg.Wait()
}

func createRecvCases(rchs []chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range rchs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	return cases
}

func Test_print_1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(300)
	catChan := make(chan struct{}, 1)
	dogChan := make(chan struct{}, 1)
	fishChan := make(chan struct{}, 1)
	for i := 0; i < 100; i++ {
		go cat1(catChan, dogChan, &wg)
		go dog1(dogChan, fishChan, &wg)
		go fish1(fishChan, catChan, &wg)
	}
	catChan <- struct{}{}
	wg.Wait()
}

func cat1(catChan, dogChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-catChan
	fmt.Println("cat")
	dogChan <- struct{}{}
}

func dog1(dogChan, fishChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-dogChan
	fmt.Println("dog")
	fishChan <- struct{}{}
}

func fish1(fishChan, catChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-fishChan
	fmt.Println("fish")
	catChan <- struct{}{}
}

func Test_print_2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	catChan := make(chan struct{})
	dogChan := make(chan struct{})
	fishChan := make(chan struct{})
	var counter uint64
	go cat2(&wg, counter, catChan, dogChan)
	go dog2(&wg, counter, dogChan, fishChan)
	go fish2(&wg, counter, fishChan, catChan)
	catChan <- struct{}{}
	wg.Wait()
}

func cat2(wg *sync.WaitGroup, counter uint64, catChan, dogChan chan struct{}) {
	for {
		if counter == 100 {
			wg.Done()
		}
		<-catChan
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogChan <- struct{}{}
	}
}

func dog2(wg *sync.WaitGroup, counter uint64, dogChan, fishChan chan struct{}) {
	for {
		if counter == 100 {
			wg.Done()
		}
		<-dogChan
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishChan <- struct{}{}
	}
}

func fish2(wg *sync.WaitGroup, counter uint64, fishChan, catChan chan struct{}) {
	for {
		if counter == 100 {
			wg.Done()
			return
		}
		<-fishChan
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		catChan <- struct{}{}
	}
}

func Test_test(t *testing.T) {
	var wg sync.WaitGroup
	catCh := make(chan struct{}, 1)
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	var counter uint64
	wg.Add(3)
	go func(catCh, dogCh chan struct{}, counter uint64) {
		for {
			if counter == 100 {
				wg.Done()
				return
			}
			<-catCh
			fmt.Println("cat")
			atomic.AddUint64(&counter, 1)
			dogCh <- struct{}{}
		}
	}(catCh, dogCh, counter)
	go func(dogCh, fishCh chan struct{}, counter uint64) {
		for {
			if counter == 100 {
				wg.Done()
				return
			}
			<-dogCh
			fmt.Println("dog")
			atomic.AddUint64(&counter, 1)
			fishCh <- struct{}{}
		}
	}(dogCh, fishCh, counter)
	go func(fishCh, catCh chan struct{}, counter uint64) {
		for {
			if counter == 100 {
				wg.Done()
				return
			}
			<-fishCh
			fmt.Println("fish")
			atomic.AddUint64(&counter, 1)
			catCh <- struct{}{}
		}
	}(fishCh, catCh, counter)
	catCh <- struct{}{}
	wg.Wait()
}
