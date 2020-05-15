// binary tree
package binary

import (
	"github.com/liunian1004/data-structrue/tree"
	"math/rand"
	"time"
)

// assert interface impl
var _ tree.Tree = (*Tree)(nil)

type Tree struct {
	Root *Node
	size int
}

type Node struct {
	Key   interface{}
	Value interface{}
	Left  *Node
	Right *Node
}

func New() *Tree {
	return &Tree{}
}

func (n *Node) insert(node *Node) {
	if n.Left == nil {
		n.Left = node
		return
	} else if n.Right == nil {
		n.Right = node
		return
	}

	rand.Seed(time.Now().UnixNano())
	lr := rand.Intn(2)
	switch lr {
	case 0:
		n.Left.insert(node)
	case 1:
		n.Right.insert(node)
	}
	return
}

func (n *Node) get(key interface{}) *Node {
	if n.Key == key {
		return n
	}

	if n.Left != nil {
		node := n.Left.get(key)
		if node != nil {
			return node
		}
	}

	if n.Right != nil {
		node := n.Right.get(key)
		if node != nil {
			return node
		}
	}

	return nil
}

func (t *Tree) Add(key, value interface{}) {
	// is key exist?

	node := &Node{
		Key: key,
		Value: value,
	}

	if t.Root == nil {
		t.Root = node
		return
	}
	t.Root.insert(node)
}

func (t *Tree) Get(key interface{}) *Node {
	return t.Root.get(key)
}

func (t *Tree) Remove() {

}

// Container a element
func (t *Tree) Container() bool {
	return false
}

func (t *Tree) Empty() bool {
	return t.size == 0
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Clear() {
	t.Root = nil
	t.size = 0
}

func (t *Tree) Values() []interface{} {
	//root := t.Root
	return nil
}