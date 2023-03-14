package DeLinkedList

import "testing"

var deLinkList = NewDeLinkList[int]()

func init() {
	deLinkList.InsertFirst(1)
	deLinkList.Insert(2, 2)
	deLinkList.Insert(3, 3)
	deLinkList.InsertLast(4)
}

func TestDeLinkList_Size(t *testing.T) {
	t.Log(deLinkList.Size())
}

func TestDeLinkList_IsEmpty(t *testing.T) {
	t.Log(deLinkList.IsEmpty())
}

func TestDeLinkList_GetValue(t *testing.T) {
	t.Log(deLinkList.GetValue())
}

func TestDeLinkList_InsertFirst(t *testing.T) {
	deLinkList.InsertFirst(10)
	t.Log(deLinkList.GetValue())
}

func TestDeLinkList_InsertLast(t *testing.T) {
	deLinkList.InsertLast(10)
	t.Log(deLinkList.GetValue())
}

func TestDeLinkList_Insert(t *testing.T) {
	deLinkList.Insert(2, 10)
	t.Log(deLinkList.GetValue())
}

func TestDeLinkList_RemoveFirst(t *testing.T) {
	deLinkList.RemoveFirst()
	t.Log(deLinkList.GetValue())
}

func TestDeLinkList_RemoveLast(t *testing.T) {
	deLinkList.RemoveLast()
	t.Log(deLinkList.GetValue())
}

func TestDeLinkList_Remove(t *testing.T) {
	deLinkList.Remove(2)
	t.Log(deLinkList.GetValue())
}

func TestDeLinkList_Contains(t *testing.T) {
	t.Log(deLinkList.Contains(2))
}

func TestDeLinkList_Find(t *testing.T) {
	t.Log(deLinkList.Find(2))
}

func TestDeLinkList_Reverse(t *testing.T) {
	deLinkList.Reverse()
	t.Log(deLinkList.GetValue())
}

func TestDeLinkList_Print(t *testing.T) {
	deLinkList.Print()
}

func TestDeLinkList_Clear(t *testing.T) {
	deLinkList.Clear()
	t.Log(deLinkList.GetValue())
}
