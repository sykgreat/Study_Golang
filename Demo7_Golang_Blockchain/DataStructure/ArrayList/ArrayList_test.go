package ArrayList

import (
	"testing"
)

var list = NewArrayList(10)

func init() {
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
}

func TestArrayList_Size(t *testing.T) {
	t.Log(list.Size())
}

func TestArrayList_Get(t *testing.T) {
	get := list.Get(1)
	t.Log(get)
}

func TestArrayList_Set(t *testing.T) {
	list.Set(1, 5)
	t.Log(list)
}

func TestArrayList_Insert(t *testing.T) {
	list.Insert(1, 5)
	t.Log(list)
}

func TestArrayList_Append(t *testing.T) {
	for i := 0; i < 10; i++ {
		list.Append(5)
	}
	t.Log(list)
}

func TestArrayList_Clear(t *testing.T) {
	list.Clear()
	t.Log(list)
}

func TestArrayList_Delete(t *testing.T) {
	list.Delete(1)
	t.Log(list)
}

func TestArrayList_String(t *testing.T) {
	t.Log(list.String())
}
