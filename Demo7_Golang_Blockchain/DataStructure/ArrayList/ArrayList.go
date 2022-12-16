package ArrayList

import (
	"errors"
	"fmt"
)

// List List接口
type List interface {
	Size() int                            // 返回数组的大小
	Get(index int) interface{}            // 返回数组中指定位置的元素
	Set(index int, newVal interface{})    // 将数组中指定位置的元素设置为新值
	Insert(index int, newVal interface{}) // 在数组中指定位置插入新元素
	Append(newVal interface{})            // 在数组末尾添加新元素
	Clear()                               // 清空数组
	Delete(index int)                     // 删除数组中指定位置的元素
	String() string                       // 返回数组的字符串表示
	Iterator() Iterator                   // 返回数组的迭代器
}

// ArrayList 数组
type ArrayList struct {
	dataStore []interface{} // 存储数据的切片
	theSize   int           // 数组的大小
}

// NewArrayList 构造函数
func NewArrayList(capSize int) *ArrayList {
	return &ArrayList{
		dataStore: make([]interface{}, 0, capSize),
		theSize:   0,
	}
}

// Size 返回数组的大小
func (list *ArrayList) Size() int {
	return list.theSize
}

// Get 返回数组中指定位置的元素
func (list *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= list.theSize {
		panic(errors.New("index out of range"))
	}
	return list.dataStore[index]
}

// Set 将数组中指定位置的元素设置为新值
func (list *ArrayList) Set(index int, newVal interface{}) {
	if index < 0 || index >= list.theSize {
		panic(errors.New("index out of range"))
	}
	list.dataStore[index] = newVal
}

// Insert 在数组中指定位置插入新元素
func (list *ArrayList) Insert(index int, newVal interface{}) {
	if index < 0 || index > list.theSize {
		panic(errors.New("index out of range"))
	}
	list.dataStore = append(list.dataStore, 0)
	copy(list.dataStore[index+1:], list.dataStore[index:])
	list.dataStore[index] = newVal
	list.theSize++
}

// Append 在数组末尾添加新元素
func (list *ArrayList) Append(newVal interface{}) {
	list.dataStore = append(list.dataStore, newVal)
	list.theSize++
}

// Clear 清空数组
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
}

func (list *ArrayList) Delete(index int) {
	if index < 0 || index >= list.theSize {
		panic(errors.New("index out of range"))
	}
	copy(list.dataStore[index:], list.dataStore[index+1:])
	list.dataStore = list.dataStore[:list.theSize-1]
	list.theSize--
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
