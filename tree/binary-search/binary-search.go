package binary_search

import (
	"fmt"
	ds "github.com/liunian1004/data-structrue"
)

type Tree struct {
	Root       *Node
	size       int
	Comparable ds.Comparable
}

type Node struct {
	Key   interface{}
	Value interface{}
	Left  *Node
	Right *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Key)
}

func New(comparator ds.Comparable) *Tree {
	return &Tree{
		Comparable: comparator,
	}
}

// 二分找到插入的位置
func (t *Tree) Add(key, value interface{}) {
	root := t.Root
	if root == nil {
		t.Root = &Node{Key: key, Value: value}
		return
	}

	t.insert(root, key, value)
}

func (t *Tree) insert(node *Node, key, value interface{}) {
	switch t.Comparable(key, node.Key) {
	case -1:
		// key < node key
		if node.Left != nil {
			t.insert(node.Left, key, value)
		} else {
			node.Left = &Node{Key: key, Value: value}
		}
	case 0:
		node.Value = value
	case 1:
		// key > node key
		if node.Right != nil {
			t.insert(node.Right, key, value)
		} else {
			node.Right = &Node{Key: key, Value: value}
		}
	}
}

// Get 的逻辑基本类似 Add
func (t *Tree) Get(key interface{}) (value interface{}, find bool) {
	return t.get(t.Root, key)
}

func (t *Tree) get(node *Node, key interface{}) (value interface{}, find bool) {
	if node == nil {
		return
	}
	switch t.Comparable(key, node.Key) {
	case -1:
		return t.get(node.Left, key)
	case 0:
		return node.Value, true
	case 1:
		return t.get(node.Right, key)
	}
	return
}

// leetcode 450
// 删除比较复杂：
// 1. 先找到对应元素的位置
// 如果 left 或 rigth 其中一个是 nil 直接把非 nil 的那个节点复制过来；
// 如果子节点是叶子节点，没有子树，删除自己？ 空间换时间，存下要删除节点的父节点：
//     if is leaf node -> find parent node -> set left or right to nil
// 如果左右都不是 nil，就要重构二叉树了，递归把 Right 节点提升（因为 right 必然大于 left）：
// 应该有很多不同的方法重构 BST
func (t *Tree) Remove(key interface{}) {

}

// 前序遍历
func (t *Tree) PreOrderTraverse() []interface{} {
	var keys []interface{}
	if t.Root == nil {
		preOrder(t.Root, keys)
	}
	return keys
}

func preOrder(node *Node, keys []interface{}) {
	if node == nil {
		return
	}
	keys = append(keys, node.Key)
	preOrder(node.Left, keys)
	preOrder(node.Right, keys)
}

// 中序遍历
func (t *Tree) InOrderTraverse() []interface{} {
	var keys []interface{}
	if t.Root == nil {
		inOrder(t.Root, keys)
	}
	return keys
}

func inOrder(node *Node, keys []interface{}) {
	if node == nil {
		return
	}
	inOrder(node.Left, keys)
	keys = append(keys, node.Key)
	inOrder(node.Right, keys)
}

// 后续遍历
func (t *Tree) PostOrderTraverse() []interface{} {
	var keys []interface{}
	if t.Root == nil {
		postOrder(t.Root, keys)
	}
	return keys
}

func postOrder(node *Node, keys []interface{}) {
	if node == nil {
		return
	}
	postOrder(node.Left, keys)
	keys = append(keys, node.Key)
	postOrder(node.Right, keys)
}

// 深度优先搜索
func (t *Tree) BFS() (values []interface{}) {
	var Queue []*Node
	Queue = append(Queue, t.Root)
	for len(Queue) != 0 {
		var newQueue []*Node
		for _, node := range Queue {
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
			values = append(values, node.Value)
		}
		Queue = newQueue
	}
	return
}

func (t *Tree) Paint() string {
	str := "BST:\n"
	if t.Root != nil {
		paint(t.Root, "", true, &str)
	}
	return str
}

func paint(node *Node, prefix string, isTail bool, str *string) {
	// 先往右边遍历到底部
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		paint(node.Right, newPrefix, false, str)
	}

	// right tree over
	*str += prefix

	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	// root node
	*str += node.String() + "\n"

	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		paint(node.Left, newPrefix, true, str)
	}
}
