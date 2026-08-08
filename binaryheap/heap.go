package binaryheap

// parent := (i - 1) / 2
// leftChild := (2 * i) + 1
// rightChild := (2 * i) + 2

func New[T any](lessThan func(a, b T) bool) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		lessThan: lessThan,
	}
}

func NewCapacity[T any](lessThan func(a, b T) bool, capacity int) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		lessThan: lessThan,
		data:     make([]T, 0, capacity),
	}
}

type BinaryHeap[T any] struct {
	data     []T
	lessThan func(a, b T) bool
}

func (bh *BinaryHeap[T]) Len() int {
	return len(bh.data)
}

func (bh *BinaryHeap[T]) Clear() {
	clear(bh.data)
}

func (bh *BinaryHeap[T]) Push(item T) {
	bh.data = append(bh.data, item)
	bh.moveUp(len(bh.data) - 1)
}

func (bh *BinaryHeap[T]) Pop() (T, bool) {
	if len(bh.data) == 0 {
		var empty T
		return empty, false
	}

	smallest := bh.data[0]
	bh.data[0] = bh.data[len(bh.data)-1]
	bh.data = bh.data[:len(bh.data)-1]
	bh.moveDown(0)

	return smallest, true
}

func (bh *BinaryHeap[T]) Peek() (T, bool) {
	if len(bh.data) == 0 {
		var empty T
		return empty, false
	}

	smallest := bh.data[0]
	return smallest, true
}

func (bh *BinaryHeap[T]) moveUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2

		if !bh.lessThan(bh.data[i], bh.data[parent]) {
			break
		}

		bh.swap(parent, i)
		i = parent
	}
}

func (bh *BinaryHeap[T]) moveDown(i int) {
	for {
		smallest := i
		leftIdx := (2 * i) + 1
		rightIdx := (2 * i) + 2

		if leftIdx < len(bh.data) && bh.lessThan(bh.data[leftIdx], bh.data[i]) {
			smallest = leftIdx
		}

		if rightIdx < len(bh.data) &&
			bh.lessThan(bh.data[rightIdx], bh.data[smallest]) {
			smallest = rightIdx
		}

		if smallest == i {
			break
		}

		bh.swap(i, smallest)
		i = smallest
	}
}

func (bh *BinaryHeap[T]) swap(leftIdx, rightIdx int) {
	temp := bh.data[leftIdx]
	bh.data[leftIdx] = bh.data[rightIdx]
	bh.data[rightIdx] = temp
}
