package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

// 饿汉
//var instance *singleton = new(singleton)
//
//func GetSingleton() *singleton {
//	return instance
//}

//// 懒汉 线程安全 加锁
//var instance *singleton
//
//// 定义一个锁
//var lock sync.Mutex
//
//var initialized uint32
//
//func GetSingleton() *singleton {
//	if atomic.LoadUint32(&initialized) == 1 {
//		return instance
//	}
//
//	lock.Lock()
//	defer lock.Unlock()
//
//	if instance == nil {
//		instance = new(singleton)
//		atomic.StoreUint32(&initialized, 1)
//		return instance
//	}
//	return instance
//}

// golang 提供的 once
var instance *singleton

var once sync.Once

func GetSingleton() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("调用了单例")
}

func main() {
	s1 := GetSingleton()
	s2 := GetSingleton()
	if s1 == s2 {
		s1.SomeThing()
		s2.SomeThing()
	}
}
