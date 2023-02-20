package cache

import (
	"log"
	"sync"
	"time"
)

type MemCache struct {
	maxMemorySize       int64  // 最大内存
	maxMemorySizeString string // 最大内存
	currentMemorySize   int64  // 当前内存

	values map[string]*memCacheValue // 缓存

	locker sync.RWMutex // 读写锁

	clearExpiredItemTimeInterval time.Duration // 清理过期缓存时间间隔
}

type memCacheValue struct {
	value      interface{} // 缓存值
	expireTime time.Time   // 过期时间
	validTime  time.Duration
	size       int64 // 缓存大小
}

func NewMemCache() *MemCache {
	m := &MemCache{
		values:                       make(map[string]*memCacheValue, 0),
		clearExpiredItemTimeInterval: time.Second * 15,
	}
	go m.clearExpiredTime()
	return m
}

func (m *MemCache) SetMaxMemory(size string) bool {
	m.maxMemorySize, m.maxMemorySizeString = ParseSize(size)
	return true
}

func (m *MemCache) Set(key string, value interface{}, expire time.Duration) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	v := &memCacheValue{
		value:      value,
		expireTime: time.Now().Add(expire),
		validTime:  expire,
		size:       GetValueSize(value),
	}
	m.del(key)
	m.add(key, v)
	if m.currentMemorySize > m.maxMemorySize {
		m.del(key)
		log.Fatalln("out of memory")
	}
	return true
}

func (m *MemCache) get(key string) (*memCacheValue, bool) {
	value, ok := m.values[key]
	return value, ok
}

func (m *MemCache) del(key string) {
	temp, ok := m.get(key)
	if ok && temp != nil {
		m.currentMemorySize -= temp.size
		delete(m.values, key)
	}
}

func (m *MemCache) add(key string, value *memCacheValue) {
	m.values[key] = value
	m.currentMemorySize += value.size
}

func (m *MemCache) Get(key string) (interface{}, bool) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	mv, ok := m.get(key)
	if ok {
		if mv.validTime != 0 && mv.expireTime.Before(time.Now()) {
			m.del(key)
			return nil, false
		}
		return mv.value, true
	}
	return nil, false
}

func (m *MemCache) Del(key string) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.del(key)
	return true
}

func (m *MemCache) Exists(key string) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()
	_, ok := m.Get(key)
	return ok
}

func (m *MemCache) Flush() bool {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.values = make(map[string]*memCacheValue, 0)
	m.currentMemorySize = 0
	return true
}

func (m *MemCache) Keys() int64 {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return int64(len(m.values))
}

func (m *MemCache) clearExpiredTime() {
	timeTicker := time.NewTicker(m.clearExpiredItemTimeInterval)
	defer timeTicker.Stop()

	for {
		select {
		case <-timeTicker.C:
			for key, item := range m.values {
				if item.validTime != 0 && time.Now().After(item.expireTime) {
					m.locker.Lock()
					m.del(key)
					m.locker.Unlock()
				}
			}
			m.Flush()
		}
	}
}
