package RedBlackTree

import "testing"

var tree = NewRedBlackTree[int]()

func init() {
	tree.Add(1, 1)
	tree.Add(2, 2)
	tree.Add(3, 3)
	tree.Add(4, 4)
}

func TestRedBlackTree_IsEmpty(t *testing.T) {
	isEmpty := tree.IsEmpty()
	t.Log(isEmpty)
}

func TestRedBlackTree_Add(t *testing.T) {
	tree.Add(5, 5)
}

func TestRedBlackTree_GetMin(t *testing.T) {
	min := tree.GetMin()
	t.Log(min)
}

func TestRedBlackTree_GetMax(t *testing.T) {
	max := tree.GetMax()
	t.Log(max)
}

func TestRedBlackTree_SearchValueForKey(t *testing.T) {
	node := tree.SearchValueForKey(3)
	t.Log(node)
}

func TestRedBlackTree_PrintPreOrder(t *testing.T) {
	tree.PrintPreOrder()
}

func TestRedBlackTree_PrintInOrder(t *testing.T) {
	tree.PrintInOrder()
}

func TestRedBlackTree_PrintPostOrder(t *testing.T) {
	tree.PrintPostOrder()
}

func TestRedBlackTree_Clean(t *testing.T) {
	tree.Clean()
}

func TestRedBlackTree_String(t *testing.T) {
	t.Log(tree)
}
