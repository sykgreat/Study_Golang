package HashMap

import (
	"unsafe"
)

type KV[K comparable, V any] struct {
	Key   K //键
	Value V // 值
}

type MyLinkNode[K comparable, V any] interface {
	InsertLast(data KV[K, V]) int //在链表尾部插入数据
}

// LinkNode 链表结构
type LinkNode[K comparable, V any] struct {
	Data KV[K, V]        //节点数据
	Next *LinkNode[K, V] //下一个节点
}

// NewLinkNode 创建链表
func NewLinkNode[K comparable, V any]() *LinkNode[K, V] {
	return &LinkNode[K, V]{
		Data: KV[K, V]{},
		Next: nil,
	}
}

// InsertLast 在链表尾部插入数据
func (ln *LinkNode[K, V]) InsertLast(data KV[K, V]) int {
	var count = 0
	//找到当前链表尾节点
	tail := ln
	for {
		count += 1
		if tail.Next == nil {
			break
		} else {
			tail = tail.Next
		}
	}

	var newNode = &LinkNode[K, V]{Data: data}
	tail.Next = newNode

	return count + 1
}

const BucketCount = 100 //桶的数量

type MyHashMap[K comparable, V any] interface {
	AddKeyValue(key K, value V) //添加键值对
	GetValueForKey(key K) V     //根据key获取value
	RemoverValueForKey(key K)   //根据key删除value
}

// HashMap 哈希表
type HashMap[K comparable, V any] struct {
	Buckets [BucketCount]*LinkNode[K, V] //桶
}

// CreateHashMap 创建HashMap
func CreateHashMap[K comparable, V any]() *HashMap[K, V] {
	myMap := &HashMap[K, V]{}
	//为每个元素添加一个链表对象
	for i := 0; i < BucketCount; i++ {
		myMap.Buckets[i] = NewLinkNode[K, V]()
	}
	return myMap
}

func (hm *HashMap[K, V]) AddKeyValue(key K, value V) {
	hashCode := memhash(unsafe.Pointer(&key), 1, 36) //1.根据key计算哈希值
	bucketIndex := hashCode % BucketCount            //2.根据哈希值计算出桶的索引
	linkNode := hm.Buckets[bucketIndex]              //3.根据桶的索引找到对应的链表

	//4.判断链表中是否已经存在该key
	if linkNode.Next == nil {
		linkNode.Data.Key = key
		linkNode.Data.Value = value
	} else {
		linkNode.InsertLast(KV[K, V]{Key: key, Value: value})
	}
}

func (hm *HashMap[K, V]) GetValueForKey(key K) (v V) {
	hashCode := memhash(unsafe.Pointer(&key), 1, 36) //1.根据key计算哈希值
	bucketIndex := hashCode % BucketCount            //2.根据哈希值计算出桶的索引
	linkNode := hm.Buckets[bucketIndex]              //3.根据桶的索引找到对应的链表

	//遍历找到key对应的节点
	head := linkNode
	for {
		if head.Data.Key == key {
			v = head.Data.Value
			break
		} else {
			if head.Next == nil {
				break
			} else {
				head = head.Next
			}
		}
	}

	return v
}

func (hm *HashMap[K, V]) RemoverValueForKey(key K) {
	hashCode := memhash(unsafe.Pointer(&key), 1, 36) //1.根据key计算哈希值
	bucketIndex := hashCode % BucketCount            //2.根据哈希值计算出桶的索引
	linkNode := hm.Buckets[bucketIndex]              //3.根据桶的索引找到对应的链表
	for {
		if linkNode.Next == nil {
			break
		} else {
			if linkNode.Next.Data.Key == key {
				linkNode.Next = linkNode.Next.Next
				break
			}
			linkNode = linkNode.Next
		}
	}
}
