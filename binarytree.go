package ds

import "slices"

type BinaryTree[T any] struct {
	data []T
	cmp  func(a, b T) int
}

func NewBinaryTree[T any](cmp func(a, b T) int) *BinaryTree[T] {
	return &BinaryTree[T]{
		cmp: cmp,
	}
}

func (t *BinaryTree[T]) lowerBound(item T) int {
	l := 0
	r := len(t.data)

	for l < r {
		m := l + (r-l)/2

		if t.cmp(t.data[m], item) < 0 {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func (t *BinaryTree[T]) Exists(item T) bool {
	if len(t.data) == 0 {
		return false
	}

	idx := t.lowerBound(item)

	if idx >= len(t.data) {
		return false
	}

	if t.cmp(t.data[idx], item) == 0 {
		return true
	}

	return false
}

func (t *BinaryTree[T]) Insert(item T) {
	if len(t.data) == 0 {
		t.data = append(t.data, item)
		return
	}

	idx := t.lowerBound(item)

	if idx >= len(t.data) {
		t.data = slices.Insert(t.data, idx, item)
		return
	}

	if t.cmp(t.data[idx], item) == 0 {
		t.data[idx] = item
		return
	}

	t.data = slices.Insert(t.data, idx, item)
}

func (t *BinaryTree[T]) Remove(item T) bool {
	if len(t.data) == 0 {
		return false
	}

	idx := t.lowerBound(item)

	if idx >= len(t.data) {
		return false
	}

	if t.cmp(t.data[idx], item) == 0 {
		t.data = append(t.data[:idx], t.data[idx+1:]...)
		return true
	}

	return false
}
