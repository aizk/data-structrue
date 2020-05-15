package binary_search

import (
	"fmt"
	data_structrue "github.com/liunian1004/data-structrue"
	"testing"
)

func TestTree_Paint(t *testing.T) {
	// BST:
	//│   ┌── 8
	//└── 6
	//    │       ┌── 5
	//    │       │   └── 3
	//    │   ┌── 2
	//    └── 1
	tree := New(data_structrue.IntComparable)
	tree.Add(6, 6)
	tree.Add(1, 1)
	tree.Add(2, 2)
	tree.Add(5, 5)
	tree.Add(8, 8)
	tree.Add(3, 3)
	tree.Add(3, 3)
	fmt.Println(tree.Paint())

	value, find := tree.Get(8)
	if !find {
		t.Error("not find 8")
	}
	if value.(int) != 8 {
		t.Error("value asset 8")
	}
}
