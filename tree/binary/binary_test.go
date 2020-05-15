package binary

import "testing"

func TestTree_Add(t *testing.T) {
	tree := New()
	tree.Add(1, 1)
	tree.Add(2, 2)
	tree.Add(3, 3)
	tree.Add(4, 4)
	tree.Add(5, 5)

	node := tree.Get(4)
	if node.Key != 4 {
		t.Error("Get(4) error")
	}
}
