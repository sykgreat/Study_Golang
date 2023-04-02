package __sync

import (
	"sync"
	"testing"
)

func Test_Mutex(t *testing.T) {
	type safeMap struct {
		data map[string]int
		sync.Mutex
	}

	sm := safeMap{
		data: make(map[string]int),
	}

	list := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "a"}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.Lock()
			defer sm.Unlock()
			for _, v := range list {
				if _, ok := sm.data[v]; ok {
					sm.data[v]++
				} else {
					sm.data[v] = 1
				}
			}
		}()
	}
	wg.Wait()
	t.Log(sm.data)
}

type Cache struct {
	data map[int]int
	sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[int]int),
	}
}

func (c *Cache) Get(key int) (int, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache) Set(key int, value int) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

func Test_RWMutex(t *testing.T) {
	cache := NewCache()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			cache.Set(i, i*i)
		}
	}()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			get, b := cache.Get(i)
			if b {
				t.Logf("key: %d, value: %d", i, get)
			}
		}(i)
	}
}
