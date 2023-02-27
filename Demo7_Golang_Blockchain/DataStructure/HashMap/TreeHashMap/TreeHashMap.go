package TreeHashMap

import (
	HashCode "Study_Golang/Demo7_Golang_Blockchain/DataStructure/HashMap"
	"Study_Golang/Demo7_Golang_Blockchain/DataStructure/Tree/RedBlackTree"
	"golang.org/x/exp/constraints"
	"reflect"
)

const BucketCount = 100 //桶的数量

type TreeHashMap[K constraints.Ordered, V any] struct {
	Buckets [BucketCount]*RedBlackTree.RedBlackTree[K, V] //桶
}

// NewTreeHashMap 创建TreeHashMap
func NewTreeHashMap[K constraints.Ordered, V any]() *TreeHashMap[K, V] {
	var buckets [BucketCount]*RedBlackTree.RedBlackTree[K, V]
	for i := 0; i < BucketCount; i++ { //初始化桶
		buckets[i] = RedBlackTree.NewRedBlackTree[K, V]()
	}
	return &TreeHashMap[K, V]{
		Buckets: buckets,
	}
}

// Add 添加键值对
func (thm *TreeHashMap[K, V]) Add(key K, value V) {
	hashCode := HashCode.Sha(reflect.ValueOf(key).String()) //1.计算key的hash值
	bucketIndex := hashCode % BucketCount                   //2.根据哈希值计算出桶的索引
	tree := thm.Buckets[bucketIndex]                        //3.根据桶的索引找到对应的链表

	tree.Insert(key, value)
}

// Get 根据key获取value
func (thm *TreeHashMap[K, V]) Get(key K) any {
	hashCode := HashCode.Sha(reflect.ValueOf(key).String()) //1.计算key的hash值
	bucketIndex := hashCode % BucketCount                   //2.根据哈希值计算出桶的索引
	tree := thm.Buckets[bucketIndex]                        //3.根据桶的索引找到对应的链表

	return tree.GetValue(key)
}

// Remove 根据key删除value
func (thm *TreeHashMap[K, V]) Remove(key K) bool {
	hashCode := HashCode.Sha(reflect.ValueOf(key).String()) //1.计算key的hash值
	bucketIndex := hashCode % BucketCount                   //2.根据哈希值计算出桶的索引
	tree := thm.Buckets[bucketIndex]                        //3.根据桶的索引找到对应的链表

	return tree.Remove(key)
}
