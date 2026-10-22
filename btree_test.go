package ds

import (
	"cmp"
	"testing"
)

func TestBTree(t *testing.T) {
	tree := NewBTree[string, int](cmp.Compare)

	err := tree.Insert("a", 1)
	if err != nil {
		t.Log(err)
	}

	_, ok := tree.Remove("a")
	if ok {
	}
}
