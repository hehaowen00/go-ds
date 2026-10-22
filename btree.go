package ds

const order int = 8
const half = order / 2

type BTree[K any, V any] struct {
	data *btreenode[K, V]
	cmp  func(a, b K) int
}

type btreenode[K any, V any] struct {
	data     [order]K
	children []*btreenode[K, V]
}

func NewBTree[K any, V any](cmp func(a, b K) int) *BTree[K, V] {
	return &BTree[K, V]{
		cmp: cmp,
	}
}

func (tree *BTree[K, V]) Insert(key K, value V) error {
	return nil
}

func (tree *BTree[K, V]) Get(key K) (V, bool) {
	var empty V
	return empty, false
}

func (tree *BTree[K, V]) Remove(key K) (V, bool) {
	var empty V
	return empty, false
}
