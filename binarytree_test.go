package ds

import (
	"cmp"
	"math/rand/v2"
	"slices"
	"testing"
)

func TestInsertKeepsSorted(t *testing.T) {
	tree := NewBinaryTree[int](cmp.Compare)

	for i := range 10 {
		tree.Insert(10 - i)
	}

	for i := range 10 {
		tree.Insert(10 + i)
	}

	if !slices.IsSortedFunc(tree.data, cmp.Compare) {
		t.Fatalf("data not sorted: %v", tree.data)
	}

	if want := 19; len(tree.data) != want {
		t.Fatalf("got %d elements, want %d: %v", len(tree.data), want, tree.data)
	}

	want := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	}

	if !slices.Equal(tree.data, want) {
		t.Fatalf("got %v, want %v", tree.data, want)
	}
}

func TestInsertDeduplicates(t *testing.T) {
	tree := NewBinaryTree[int](cmp.Compare)

	for _, v := range []int{5, 3, 8, 5, 1, 3} {
		tree.Insert(v)
	}

	if want := []int{1, 3, 5, 8}; !slices.Equal(tree.data, want) {
		t.Fatalf("got %v, want %v", tree.data, want)
	}
}

func TestInsertBoundaries(t *testing.T) {
	tree := NewBinaryTree[int](cmp.Compare)

	for _, v := range []int{2, 4, 6} {
		tree.Insert(v)
	}

	tree.Insert(0)
	tree.Insert(8)
	tree.Insert(6)

	if want := []int{0, 2, 4, 6, 8}; !slices.Equal(tree.data, want) {
		t.Fatalf("got %v, want %v", tree.data, want)
	}
}

func TestInsertStructOrdering(t *testing.T) {
	type Rank struct {
		Name  string
		Score int
	}

	rankCmp := func(a, b Rank) int {
		if ord := cmp.Compare(a.Score, b.Score); ord != 0 {
			return ord
		}

		return cmp.Compare(a.Name, b.Name)
	}

	tree := NewBinaryTree(rankCmp)
	tree.Insert(Rank{Name: "bob", Score: 2})
	tree.Insert(Rank{Name: "alice", Score: 1})
	tree.Insert(Rank{Name: "bob", Score: 2})
	tree.Insert(Rank{Name: "carol", Score: 1})

	want := []Rank{
		{Name: "alice", Score: 1},
		{Name: "carol", Score: 1},
		{Name: "bob", Score: 2},
	}
	if !slices.Equal(tree.data, want) {
		t.Fatalf("got %v, want %v", tree.data, want)
	}
}

func TestInsertRandomized(t *testing.T) {
	r := rand.New(rand.NewPCG(42, 7))
	const n = 1000

	tree := NewBinaryTree[int](cmp.Compare)
	seen := map[int]bool{}
	for range n {
		v := r.IntN(1000)
		tree.Insert(v)
		seen[v] = true
	}

	if !slices.IsSortedFunc(tree.data, cmp.Compare) {
		t.Fatalf("data not sorted: %v", tree.data)
	}

	if got := len(tree.data); got != len(seen) {
		t.Fatalf(
			"got %d unique elements, want %d (data: %v)", got, len(seen), tree.data,
		)
	}

	if !slices.Equal(tree.data, sortedKeys(seen)) {
		t.Fatalf("tree data does not match inserted keys: %v", tree.data)
	}
}

func TestEmptyTree(t *testing.T) {
	tree := NewBinaryTree[int](cmp.Compare)
	if len(tree.data) != 0 {
		t.Fatalf("new tree should be empty, got %v", tree.data)
	}
}

func sortedKeys(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	return keys
}
