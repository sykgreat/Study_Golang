package cache_server

import (
	"Study_Golang/Interview/memCache/cache"
	"time"
)

type CacheServer struct {
	memCache cache.Cache
}

func NewMemCache() *CacheServer {
	return &CacheServer{
		memCache: cache.NewMemCache(),
	}
}

func (cs *CacheServer) SetMaxMemory(size string) bool {
	return cs.memCache.SetMaxMemory(size)
}
func (cs *CacheServer) Set(key string, value interface{}, expire ...time.Duration) bool {
	expireTime := time.Second * 0
	if len(expire) > 0 {
		expireTime = expire[0]
	}
	return cs.memCache.Set(key, value, expireTime)
}
func (cs *CacheServer) Get(key string) (interface{}, bool) {
	return cs.memCache.Get(key)
}
func (cs *CacheServer) Del(key string) bool {
	return cs.memCache.Del(key)
}
func (cs *CacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}
func (cs *CacheServer) Flush() bool {
	return cs.memCache.Flush()
}
func (cs *CacheServer) Keys() int64 {
	return cs.memCache.Keys()
}
