package RedBlackTree

import "testing"

var tree = NewRedBlackTree[int]()

func init() {
	tree.Add(1, 1)
	tree.Add(2, 2)
	tree.Add(3, 3)
	tree.Add(4, 4)
}

func TestRedBlackTree_Add(t *testing.T) {
	tree.Add(5, 5)
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

func TestRedBlackTree_PrintTree(t *testing.T) {
	tree.PrintTree()
}
