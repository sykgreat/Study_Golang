package cache

import "time"

type Cache interface {
	SetMaxMemory(size string) bool                                // 设置最大内存
	Set(key string, value interface{}, expire time.Duration) bool // 设置缓存
	Get(key string) (interface{}, bool)                           // 获取缓存
	Del(key string) bool                                          // 删除缓存
	Exists(key string) bool                                       // 判断缓存是否存在
	Flush() bool                                                  // 清空缓存
	Keys() int64                                                  // 获取缓存数量
}
