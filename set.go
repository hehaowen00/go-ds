package ds

import "cmp"

type BTreeSet[K cmp.Ordered] struct {
}

type HashSet[K cmp.Ordered] struct {
	data map[K]struct{}
}

func NewHashSet[K cmp.Ordered]() HashSet[K] {
	return HashSet[K]{
		data: map[K]struct{}{},
	}
}

func (set *HashSet[K]) Get(key K) bool {
	_, ok := set.data[key]
	return ok
}

func (set *HashSet[K]) Insert(key K) {
	set.data[key] = struct{}{}
}

func (set *HashSet[K]) Remove(key K) bool {
	_, ok := set.data[key]
	delete(set.data, key)
	return ok
}
