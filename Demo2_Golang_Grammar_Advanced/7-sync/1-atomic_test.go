package __sync

import (
	"sync"
	"sync/atomic"
	"testing"
)

func Test_Atomic_Foundation(t *testing.T) {
	var a int64 = 10
	// 原子操作
	// 原子操作是指不会被线程调度机制打断的操作，即使是多核CPU，原子操作也是连续进行的，不会出现被打断的情况。
	// 原子操作的语义是不可分割的，要么执行成功，要么执行失败，不会出现执行了一半的情况。
	// 原子操作的执行速度非常快，通常比加锁的执行速度还要快。

	// 原子存储
	// 存储变量的值。
	atomic.StoreInt64(&a, 20)

	// 原子加载
	// 加载变量的值，并返回。
	t.Log("获取变量的值: ", atomic.LoadInt64(&a))

	// 原子累加
	// 累加变量的值，并返回累加后的值。
	t.Log("在原有基础上累加: ", atomic.AddInt64(&a, 10))

	// 原子交换
	// 交换变量的值，并返回原有的值。
	t.Log("原子交换: 交换变量的值，并返回原有的值: ", atomic.SwapInt64(&a, 40))
	t.Log("获取交换后变量的值: ", atomic.LoadInt64(&a))

	// 原子比较并交换
	// 如果变量的值等于old，则将变量的值修改为new，否则不做任何操作。
	// 返回值表示是否修改成功。
	t.Log("原子比较并交换: ", atomic.CompareAndSwapInt64(&a, 40, 50))
	t.Log("获取变量的值: ", atomic.LoadInt64(&a))
}

func Test_Concurrent_1(t *testing.T) {
	var a uint64
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a += 1
		}()
	}
	wg.Wait()
	t.Log(a)
}

func Test_Locker_1(t *testing.T) {
	var a uint64
	wg := sync.WaitGroup{}
	locker := sync.Mutex{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(locker *sync.Mutex) {
			defer wg.Done()
			locker.Lock()
			a += 1
			locker.Unlock()
		}(&locker)
	}
	wg.Wait()
	t.Log(a)
}

func Test_Atomic_1(t *testing.T) {
	var a uint64
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&a, 1)
		}()
	}
	wg.Wait()
	t.Log(a)
}

func Test_Concurrent_2(t *testing.T) {
	list := []string{"a", "b", "c", "d"}
	mp := make(map[string]int) // 普通的map
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, v := range list {
				_, ok := mp[v]
				if !ok {
					mp[v] = 0
				}
				mp[v] += 1
			}
		}()
	}
	wg.Wait()
	t.Log(mp)
}

func Test_Atomic_2_Bug_1(t *testing.T) {
	list := []string{"a", "b", "c", "d"}
	mp := atomic.Value{}
	mp.Store(make(map[string]int)) // 原子的map
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m := mp.Load().(map[string]int)
			for _, v := range list {
				_, ok := m[v]
				if !ok {
					m[v] = 0
				}
				m[v] += 1
			}
			mp.Store(m)
		}()
	}
	wg.Wait()
	t.Log(mp.Load())
}

func Test_Atomic_2_BUg_2(t *testing.T) {
	list := []string{"a", "b", "c", "d"}
	mp := atomic.Value{}
	mp.Store(make(map[string]int)) // 原子的map
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m := mp.Load().(map[string]int) // 原子操作

			// 普通操作
			m1 := make(map[string]int)
			for k, v := range m {
				m1[k] = v
			}
			for _, v := range list {
				_, ok := m1[v]
				if !ok {
					m1[v] = 0
				}
				m1[v] += 1
			}

			mp.Store(m1) // 原子操作
		}()
	}
	wg.Wait()
	t.Log(mp.Load())
}

func Test_Atomic_2(t *testing.T) {
	list := []string{"a", "b", "c", "d"}
	mp := atomic.Value{}
	mp.Store(new(map[string]int)) // 原子的map
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		atomicLabel:
			m := mp.Load().(*map[string]int) // 原子操作

			// 普通操作
			m1 := make(map[string]int)
			for k, v := range *m {
				m1[k] = v
			}
			for _, v := range list {
				_, ok := m1[v]
				if !ok {
					m1[v] = 0
				}
				m1[v] += 1
			}

			swapped := mp.CompareAndSwap(m, &m1) // 原子操作
			if !swapped {
				t.Log("交换失败")
				t.Log("重新执行, 交换逻辑")
				goto atomicLabel
			}
		}()
	}
	wg.Wait()
	t.Log(mp.Load())
}
