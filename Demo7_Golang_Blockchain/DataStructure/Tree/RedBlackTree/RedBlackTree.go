package RedBlackTree

import (
	"Study_Golang/Demo7_Golang_Blockchain/utils"
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

const (
	RED   = true  // 红色
	BLACK = false // 黑色
)

// RedBlackTreeInterface 红黑树接口
type RedBlackTreeInterface[K constraints.Ordered, V any] interface {
	GetSize() int                         // 获取红黑树的大小
	IsEmpty() bool                        // 红黑树是否为空
	Insert(key K, value V)                // 向红黑树中添加元素
	GetValue(key K) V                     // 获取key对应的value
	Search(key K) *RedBlackTreeNode[K, V] // 在红黑树中查找元素
	Remove(key K) bool                    // 从红黑树中删除元素
	InOrderPrint()                        // 中序遍历红黑树
	GetTreeDepth() int                    // 获取红黑树的深度
	Show()                                // 打印红黑树
}

// RedBlackTree 红黑树
type RedBlackTree[K constraints.Ordered, V any] struct {
	root *RedBlackTreeNode[K, V] // 根节点
	size int                     // 红黑树的大小
}

// NewRedBlackTree 创建红黑树
func NewRedBlackTree[K constraints.Ordered, V any]() *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{
		root: nil,
		size: 0,
	}
}

// GetSize 获取红黑树的大小
func (rbt *RedBlackTree[K, V]) GetSize() int {
	return rbt.size
}

// IsEmpty 红黑树是否为空
func (rbt *RedBlackTree[K, V]) IsEmpty() bool {
	return rbt.size == 0
}

// Insert 向红黑树中添加元素
func (rbt *RedBlackTree[K, V]) Insert(key K, value V) {
	node := NewRedBlackTreeNode(key, value)
	var parent *RedBlackTreeNode[K, V] = nil
	current := rbt.root

	for current != nil {
		parent = current
		if key < current.key {
			current = current.left
		} else if key > current.key {
			current = current.right
		} else {
			current.value = value
			return
		}
	}

	if parent != nil {
		node.parent = parent
		if node.key < parent.key { // 小于父节点
			parent.left = node
		} else { // 大于父节点
			parent.right = node
		}
	} else {
		rbt.root = node
	}

	rbt.size++
	node.repairRedBlackTrees(rbt)
}

// GetValue 获取key对应的value
func (rbt *RedBlackTree[K, V]) GetValue(key K) (v V) {
	node := rbt.root.search(key)
	if node != nil {
		v = node.value
	}
	return v
}

// Search 在红黑树中查找元素
func (rbt *RedBlackTree[K, V]) Search(key K) *RedBlackTreeNode[K, V] {
	return rbt.root.search(key)
}

// Remove 从红黑树中删除元素
func (rbt *RedBlackTree[K, V]) Remove(key K) (result bool) {
	result = rbt.root.remove(key, rbt)
	if result {
		rbt.size--
	}
	return result
}

// InOrderPrint 中序遍历红黑树
func (rbt *RedBlackTree[K, V]) InOrderPrint() {
	rbt.root.inOrderPrint()
}

// GetTreeDepth 获取红黑树的深度
func (rbt *RedBlackTree[K, V]) GetTreeDepth() int {
	return rbt.root.getTreeDepth()
}

// Show 打印红黑树
func (rbt *RedBlackTree[K, V]) Show() {
	if rbt.root == nil {
		return
	}

	depth := rbt.root.getTreeDepth()

	arrayHeight := depth*2 - 1
	arrayWidth := (2<<(depth-2))*3 + 1
	res := make([][]string, arrayHeight)
	for i := 0; i < arrayHeight; i++ {
		res[i] = make([]string, arrayWidth)
		for j := 0; j < arrayWidth; j++ {
			res[i][j] = " "
		}
	}

	writeArray(rbt.root, 0, arrayWidth/2, res, depth)

	for _, line := range res {
		sb := make([]string, 0)
		for j := 0; j < len(line); j++ {
			sb = append(sb, line[j])
			if len(line[j]) > 1 && j < len(line)-1 {
				if len(line[j]) > 4 {
					j += 2
				} else {
					j += len(line[j]) - 1
				}
			}

		}
		fmt.Println(strings.Join(sb, ""))
	}
}

func writeArray[K constraints.Ordered, V any](currNode *RedBlackTreeNode[K, V], rowIndex int, columnIndex int, res [][]string, treeDepth int) {
	if currNode == nil {
		return
	}

	key, _ := utils.ToStringE(currNode.key)
	//value, _ := utils.ToStringE(currNode.value)
	color := "R"
	if !currNode.color {
		color = "B"
	}

	//res[rowIndex][columnIndex] = key + "-" + value + "-" + (color) + " "
	res[rowIndex][columnIndex] = key + "-" + (color) + ""

	currLevel := (rowIndex + 1) / 2
	if currLevel == treeDepth {
		return
	}

	gap := treeDepth - currLevel - 1

	if currNode.left != nil {
		res[rowIndex+1][columnIndex-gap] = " /"
		writeArray(currNode.left, rowIndex+2, columnIndex-gap*2, res, treeDepth)
	}

	if currNode.right != nil {
		res[rowIndex+1][columnIndex+gap] = " \\"
		writeArray(currNode.right, rowIndex+2, columnIndex+gap*2, res, treeDepth)
	}
}

// RedBlackTreeNodeInterface 红黑树节点接口
type RedBlackTreeNodeInterface[K constraints.Ordered, V any] interface {
	leftRotate(*RedBlackTree[K, V])          // 左旋转
	rightRotate(*RedBlackTree[K, V])         // 右旋转
	repairRedBlackTrees(*RedBlackTree[K, V]) // 修复红黑树
	search(K) *RedBlackTreeNode[K, V]        // 查找
	remove(K, *RedBlackTree[K, V]) bool      // 删除
	getMinNode() *RedBlackTreeNode[K, V]     // 获取最小节点
	inOrderPrint()                           // 中序遍历
	getTreeDepth() int                       // 获取红黑树的深度
}

type RedBlackTreeNode[K constraints.Ordered, V any] struct {
	key   K
	value V
	color bool // 颜色

	parent *RedBlackTreeNode[K, V] // 父节点
	left   *RedBlackTreeNode[K, V] // 左子节点
	right  *RedBlackTreeNode[K, V] // 右子节点
}

// NewRedBlackTreeNode 创建红黑树节点
func NewRedBlackTreeNode[K constraints.Ordered, V any](key K, value V) *RedBlackTreeNode[K, V] {
	return &RedBlackTreeNode[K, V]{
		key:   key,
		value: value,
		color: RED,
	}
}

/**
* leftRotate 左旋转
* 1      p                             p
* 1      |                             |
* 1     rbtn                           y
* 1    /   \       ----------->      /   \
* 1   lx    y                      rbtn  ry
* 1        /  \                    /  \
* 1       ly  ry                  lx  ly
* 1.将y的左子节点的父节点更新为rbtn, 并将rbtn的右子节点指向y的左子节点（ly)
* 2.当rbtn的父节点(不为空时)﹐更新y的父节点为rbtn的父节点, 并将rbtn的父节点指定子树（当前rbtn的子树位置)指定为y
* 3.将rbtn的父节点更新为y, 将y的左子节点更新为node
*
* @param x 当前节点
 */
func (rbtn *RedBlackTreeNode[K, V]) leftRotate(tree *RedBlackTree[K, V]) {
	y := rbtn.right
	rbtn.right = y.left

	if y.left != nil {
		y.left.parent = rbtn
	}

	if rbtn.parent != nil {
		y.parent = rbtn.parent
		if rbtn.parent.left == rbtn {
			rbtn.parent.left = y
		} else {
			rbtn.parent.right = y
		}
	} else {
		tree.root = y
		y.parent = nil
	}

	rbtn.parent = y
	y.left = rbtn
}

/**
* rightRotate 右旋转
* 1      p                             p
* 1      |                             |
* 1     rbtn                           x
* 1    /   \       ----------->      /   \
* 1   x    ry                      lx   rbtn
* 1  /  \                               /  \
* 1 lx  rx                             rx  ry
* 1.将x的右子节点的父节点更新为rbtn, 并将rbtn的左子节点指向x的右子节点（rx)
* 2.当rbtn的父节点(不为空时)﹐更新x的父节点为rbtn的父节点, 并将rbtn的父节点指定子树（当前rbtn的子树位置)指定为x
* 3.将rbtn的父节点更新为x, 将x的右子节点更新为node
 */
func (rbtn *RedBlackTreeNode[K, V]) rightRotate(tree *RedBlackTree[K, V]) {
	x := rbtn.left
	rbtn.left = x.right

	if x.right != nil {
		x.right.parent = rbtn
	}

	if rbtn.parent != nil {
		x.parent = rbtn.parent
		if rbtn.parent.left == rbtn {
			rbtn.parent.left = x
		} else {
			rbtn.parent.right = x
		}
	} else {
		tree.root = x
		x.parent = nil
	}

	rbtn.parent = x
	x.right = rbtn
}

// repairRedBlackTrees 修复红黑树
func (rbtn *RedBlackTreeNode[K, V]) repairRedBlackTrees(tree *RedBlackTree[K, V]) {
	if rbtn.parent == nil {
		rbtn.color = BLACK
		return
	}

	parent := rbtn.parent
	var grandParent *RedBlackTreeNode[K, V] = nil
	if parent != nil {
		grandParent = parent.parent
	}

	if parent != nil && parent.color == RED {
		var uncle *RedBlackTreeNode[K, V] = nil
		if parent == grandParent.left {
			uncle = grandParent.right
		} else {
			uncle = grandParent.left
		}

		if uncle != nil && uncle.color == RED {
			parent.color = BLACK
			uncle.color = BLACK
			grandParent.color = RED
			grandParent.repairRedBlackTrees(tree)
		} else if parent == grandParent.left {
			if rbtn == parent.left {
				parent.color = BLACK
				grandParent.color = RED
				grandParent.rightRotate(tree)
			} else {
				rbtn = parent
				rbtn.leftRotate(tree)
				rbtn.repairRedBlackTrees(tree)
			}
		} else {
			if rbtn == parent.right {
				parent.color = BLACK
				grandParent.color = RED
				grandParent.leftRotate(tree)
			} else {
				rbtn = parent
				rbtn.rightRotate(tree)
				rbtn.repairRedBlackTrees(tree)
			}
		}
	}
}

// search 查找
func (rbtn *RedBlackTreeNode[K, V]) search(key K) *RedBlackTreeNode[K, V] {
	if rbtn == nil {
		return nil
	}

	if rbtn.key == key {
		return rbtn
	} else if rbtn.key < key {
		return rbtn.right.search(key)
	} else {
		return rbtn.left.search(key)
	}
}

// remove 删除
func (rbtn *RedBlackTreeNode[K, V]) remove(k K, tree *RedBlackTree[K, V]) bool {
	if rbtn == nil {
		return false
	}

	if rbtn.key < k { // 当前节点小于要删除的节点, 则在右子树中删除
		return rbtn.right.remove(k, tree)
	} else if rbtn.key > k { // 当前节点大于要删除的节点, 则在左子树中删除
		return rbtn.left.remove(k, tree)
	} else { // 当前节点等于要删除的节点, 则删除当前节点
		if rbtn.left != nil && rbtn.right != nil { // 当前节点有两个子节点
			minNode := rbtn.right.getMinNode()
			rbtn.key = minNode.key
			rbtn.value = minNode.value
			rbtn.right.remove(minNode.key, tree)
		} else if rbtn.left != nil { // 当前节点只有左子节点
			if rbtn.parent == nil { // 当前节点是根节点
				rbtn.left.parent = nil
				rbtn = rbtn.left
			} else if rbtn.parent.left == rbtn { // 当前节点是父节点的左子节点
				rbtn.parent.left = rbtn.left
				rbtn.left.parent = rbtn.parent
			} else { // 当前节点是父节点的右子节点
				rbtn.parent.right = rbtn.left
				rbtn.left.parent = rbtn.parent
			}
		} else if rbtn.right != nil { // 当前节点只有右子节点
			if rbtn.parent == nil { // 当前节点是根节点
				rbtn.right.parent = nil
				rbtn = rbtn.right
			} else if rbtn.parent.left == rbtn { // 当前节点是父节点的左子节点
				rbtn.parent.left = rbtn.right
				rbtn.right.parent = rbtn.parent
			} else { // 当前节点是父节点的右子节点
				rbtn.parent.right = rbtn.right
				rbtn.right.parent = rbtn.parent
			}
		} else { // 当前节点没有子节点
			if rbtn.parent == nil { // 当前节点是根节点
				rbtn = nil
				tree.root = nil
			} else if rbtn.parent.left == rbtn { // 当前节点是父节点的左子节点
				rbtn.parent.left = nil
			} else { // 当前节点是父节点的右子节点
				rbtn.parent.right = nil
			}
		}
	}

	return true
}

// getMinNode 获取最小节点
func (rbtn *RedBlackTreeNode[K, V]) getMinNode() *RedBlackTreeNode[K, V] {
	if rbtn.left == nil {
		return rbtn
	}

	return rbtn.left.getMinNode()
}

// inOrderPrint 中序遍历
func (rbtn *RedBlackTreeNode[K, V]) inOrderPrint() {
	if rbtn == nil {
		return
	}

	rbtn.left.inOrderPrint()
	fmt.Printf("%v ", rbtn.value)
	rbtn.right.inOrderPrint()
}

// getTreeDepth 获取红黑树的深度
func (rbtn *RedBlackTreeNode[K, V]) getTreeDepth() int {
	if rbtn == nil {
		return 0
	}

	leftDepth := rbtn.left.getTreeDepth()
	rightDepth := rbtn.right.getTreeDepth()

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
