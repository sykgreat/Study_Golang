package __sync

import (
	"sync"
	"testing"
)

func Test_Map_Foundation(t *testing.T) {
	sm := sync.Map{}

	// 设置键值对
	sm.Store("a", 1)
	sm.Store("b", 2)

	// 获取键值对
	t.Log(sm.Load("a"))
	t.Log(sm.Load("b"))

	// 通过key设置value 如果不存在则设置指定的value并返回
	// ok为true表示key存在并返回值，为false表示key不存在并设置后返回
	t.Log(sm.LoadOrStore("c", 3))
	t.Log(sm.LoadOrStore("c", 4))

	// 根据key获取value后 删除key-value
	// ok为true表示key存在 删除，为false表示key不存在
	t.Log(sm.LoadAndDelete("b"))

	// 为集合设置迭代函数 将集合中的每一个键值对顺序调用该函数 如果该函数返回false 则停止迭代
	// 为遍历集合中的所有键值对 提供了一种简单的方式
	sm.Range(func(key, value interface{}) bool {
		t.Log(key, value)
		return true
	})
}
