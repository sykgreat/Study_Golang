package RedBlackTree

import "fmt"

const (
	RED   = true  // 红色
	BLACK = false // 黑色
)

type MyNode[T any] interface {
	IsRed() bool                           // 是否为红色
	leftRotate() *Node[T]                  // 左旋转
	rightRotate() *Node[T]                 // 右旋转
	flipColors()                           // 颜色翻转
	updateRedBlackTree(isAdd int) *Node[T] // 更新红黑树
	add(key int, value T) (int, *Node[T])  // 添加节点
	printPreOrder(resp *[][]any)           // 前序遍历
	printInOrder(resp *[][]any)            // 中序遍历
	printPostOrder(resp *[][]any)          // 后序遍历
	string() string                        // 打印节点
}

// Node 节点
type Node[T any] struct {
	Key   int
	Value T

	left  *Node[T]
	right *Node[T]
	color bool
}

// NewNode 新建节点
func NewNode[T any](key int, value T) *Node[T] {
	return &Node[T]{
		Key:   key,
		Value: value,
		left:  nil,
		right: nil,
		color: RED,
	}
}

// IsRed 是否为红色
func (n *Node[T]) IsRed() bool {
	if n == nil {
		return BLACK
	}
	return n.color
}

//		   node                    x
//		  /   \     左旋转         /  \
//		 T1   x   --------->    node  T3
//			 / \               /   \
//	 		T2 T3             T1   T2
func (n *Node[T]) leftRotate() *Node[T] {
	retNode := n.right
	n.right = retNode.left

	retNode.left = n
	retNode.color = n.color
	n.color = RED
	return retNode
}

//	   node                   x
//	  /   \     右旋转       /  \
//	 x    T2   ------->   y   node
//	/ \                       /  \
//
// y  T1                     T1  T2
func (n *Node[T]) rightRotate() *Node[T] {
	retNode := n.left
	n.left = retNode.right

	retNode.right = n
	retNode.color = n.color
	n.color = RED
	return retNode
}

func (n *Node[T]) flipColors() {
	n.color = RED
	n.left.color = BLACK
	n.right.color = BLACK
}

// 更新红黑树
func (n *Node[T]) updateRedBlackTree(isAdd int) *Node[T] {
	// isAdd=0 说明没有新节点，无需维护
	if isAdd == 0 {
		return n
	}

	// 是否需要左旋转
	if n.right.IsRed() == RED && n.left.IsRed() != RED {
		n = n.leftRotate()
	}

	// 是否需要右旋转
	if n.left.IsRed() == RED && n.left.left.IsRed() == RED {
		n = n.rightRotate()
	}

	// 是否需要颜色翻转
	if n.left.IsRed() == RED && n.right.IsRed() == RED {
		n.flipColors()
	}

	return n
}

// 添加节点
func (n *Node[T]) add(key int, value T) (int, *Node[T]) {
	if n == nil {
		return 1, NewNode[T](key, value)
	}

	isAdd := 0
	if key < n.Key { // 小于当前节点，添加到左子树
		isAdd, n.left = n.left.add(key, value) // 递归添加
	} else if key > n.Key { // 大于当前节点，添加到右子树
		isAdd, n.right = n.right.add(key, value) // 递归添加
	} else { // 等于当前节点，更新value值
		n.Value = value // 对value值更新,节点数量不增加,isAdd = 0
	}

	// 维护红黑树
	n = n.updateRedBlackTree(isAdd)

	return isAdd, n
}

// 前序遍历
func (n *Node[T]) printPreOrder(resp *[][]any) {
	if n == nil {
		return
	}
	*resp = append(*resp, []any{n.Key, n.Value, n.color})
	n.left.printPreOrder(resp)
	n.right.printPreOrder(resp)
}

// 中序遍历
func (n *Node[T]) printInOrder(resp *[][]any) {
	if n == nil {
		return
	}
	n.left.printInOrder(resp)
	*resp = append(*resp, []any{n.Key, n.Value, n.color})
	n.right.printInOrder(resp)
}

// 后序遍历
func (n *Node[T]) printPostOrder(resp *[][]any) {
	if n == nil {
		return
	}
	n.left.printPostOrder(resp)
	n.right.printPostOrder(resp)
	*resp = append(*resp, []any{n.Key, n.Value, n.color})
}

// 打印节点
func (n *Node[T]) string() string {
	return fmt.Sprintf("key:%v,value:%v,color:%v", n.Key, n.Value, n.color)
}

type MyRedBlackTree[T any] interface {
	IsEmpty() bool                      // 判断树是否为空
	GetTreeSize() int                   // 获取树的节点数量
	Add(key int, value T)               // 添加节点
	GetMin() *Node[T]                   // 获取最小值
	GetMax() *Node[T]                   // 获取最大值
	SearchValueForKey(key int) *Node[T] // 根据key查找value
	PrintPreOrder()                     // 前序遍历
	PrintInOrder()                      // 中序遍历
	PrintPostOrder()                    // 后序遍历
	Clean()                             // 清空树
	String() string                     // 打印树
}

// RedBlackTree 红黑树
type RedBlackTree[T any] struct {
	root *Node[T]
	size int
}

// NewRedBlackTree 新建红黑树
func NewRedBlackTree[T any]() *RedBlackTree[T] {
	return &RedBlackTree[T]{
		root: nil,
		size: 0,
	}
}

// IsEmpty 判断树是否为空
func (rbt *RedBlackTree[T]) IsEmpty() bool {
	return rbt.size == 0
}

// GetTreeSize 获取树的节点数量
func (rbt *RedBlackTree[T]) GetTreeSize() int {
	return rbt.size
}

// Add 添加节点
func (rbt *RedBlackTree[T]) Add(key int, value T) {
	isAdd, nd := rbt.root.add(key, value)
	rbt.size += isAdd
	rbt.root = nd
	rbt.root.color = BLACK //根节点为黑色节点
}

// GetMin 获取最小值
func (rbt *RedBlackTree[T]) GetMin() (min *Node[T]) {
	min = rbt.root
	for min.left != nil {
		min = min.left
	}
	return min
}

// GetMax 获取最大值
func (rbt *RedBlackTree[T]) GetMax() (max *Node[T]) {
	max = rbt.root
	for max.right != nil {
		max = max.right
	}
	return max
}

// SearchValueForKey 根据key查找value
func (rbt *RedBlackTree[T]) SearchValueForKey(key int) (node *Node[T]) {
	node = rbt.root
	for node != nil {
		if key < node.Key {
			node = node.left
		} else if key > node.Key {
			node = node.right
		} else {
			break
		}
	}
	return node
}

// PrintPreOrder 前序遍历
func (rbt *RedBlackTree[T]) PrintPreOrder() {
	var resp [][]any
	rbt.root.printPreOrder(&resp)
	fmt.Println(resp)
}

// PrintInOrder 中序遍历
func (rbt *RedBlackTree[T]) PrintInOrder() {
	var resp [][]any
	rbt.root.printInOrder(&resp)
	fmt.Println(resp)
}

// PrintPostOrder 后序遍历
func (rbt *RedBlackTree[T]) PrintPostOrder() {
	var resp [][]any
	rbt.root.printPostOrder(&resp)
	fmt.Println(resp)
}

// Clean 清空树
func (rbt *RedBlackTree[T]) Clean() {
	rbt.root = nil
	rbt.size = 0
}

// String 打印树
func (rbt *RedBlackTree[T]) String() string {
	str := "RedBlackTree\n"
	if !rbt.IsEmpty() {
		output(rbt.root, "", true, &str)
	}
	return str
}

func output[T any](node *Node[T], prefix string, isTail bool, str *string) {
	if node.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.string() + "\n"
	if node.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.left, newPrefix, true, str)
	}
}
