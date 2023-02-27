package RedBlackTree

import "testing"

var tree = NewRedBlackTree[int, int]()

//func init() {
//	tree.Insert(1, 1)
//	tree.Insert(2, 2)
//	tree.Insert(3, 3)
//	tree.Insert(4, 4)
//	tree.Insert(5, 5)
//	tree.Insert(6, 6)
//	tree.Insert(7, 7)
//	tree.Insert(8, 8)
//	tree.Insert(9, 9)
//	tree.Insert(10, 10)
//	tree.Insert(11, 11)
//	tree.Insert(12, 12)
//	tree.Insert(13, 13)
//	tree.Insert(14, 14)
//	tree.Insert(15, 15)
//	tree.Insert(16, 16)
//	tree.Insert(17, 17)
//	tree.Insert(18, 18)
//	tree.Insert(19, 19)
//	tree.Insert(20, 20)
//}

func TestRedBlackTree_IsEmpty(t *testing.T) {
	t.Log(tree.IsEmpty())
}

func TestRedBlackTree_GetSize(t *testing.T) {
	t.Log(tree.GetSize())
}

func TestRedBlackTree_Insert(t *testing.T) {
	tree.Insert(5, 5)
}

func TestRedBlackTree_Search(t *testing.T) {
	t.Log(tree.Search(12))
}

func TestRedBlackTree_Remove(t *testing.T) {
	t.Log(tree.Remove(8))
	tree.Show()
}

func TestRedBlackTree_GetTreeDepth(t *testing.T) {
	t.Log(tree.GetTreeDepth())
	tree.Show()
}

func TestRedBlackTree_InOrderPrint(t *testing.T) {
	tree.InOrderPrint()
}

func TestRedBlackTree_Performance_Insert(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		tree.Insert(i, i)
	}
}

func TestRedBlackTree_Performance_Search(t *testing.T) {
	t.Log(tree.Search(100000))
}
