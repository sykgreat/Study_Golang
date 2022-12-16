package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool     // 判断是否有下一个元素
	Next() interface{} // 返回数组中指定位置的元素
	Remove()           // 删除当前位置的元素
	GetIndex() int     // 返回当前位置
}

type Iterable interface {
	Iterator() Iterator // 构造初始化接口 返回一个迭代器
}

// ArrayListIterator 构造指针 访问数组
type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前索引
}

func (list *ArrayList) Iterator() Iterator {
	return &ArrayListIterator{
		list:         list,
		currentIndex: 0,
	}
}

func (iterator *ArrayListIterator) HasNext() bool {
	return iterator.currentIndex < iterator.list.theSize // 判断当前索引是否小于数组大小
}

func (iterator *ArrayListIterator) Next() interface{} {
	if !iterator.HasNext() {
		panic(errors.New("no next element"))
	}
	value := iterator.list.Get(iterator.currentIndex)
	iterator.currentIndex++
	return value
}

func (iterator *ArrayListIterator) Remove() {
	iterator.currentIndex--
	iterator.list.Delete(iterator.currentIndex)
}

func (iterator *ArrayListIterator) GetIndex() int {
	return iterator.currentIndex
}
