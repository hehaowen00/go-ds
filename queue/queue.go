package queue

import "github.com/hehaowen00/go-ds/binaryheap"

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.data) == 0 {
		var empty T
		return empty, false
	}

	item := q.data[0]
	q.data = q.data[1:]
	return item, true
}

func NewPriority[T PriorityItem]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: binaryheap.New(func(a, b T) bool {
			return a.Priority() < b.Priority()
		}),
	}
}

type PriorityQueue[T PriorityItem] struct {
	heap *binaryheap.BinaryHeap[T]
}

type PriorityItem interface {
	Priority() int64
}

func (q *PriorityQueue[T]) Len() int {
	return q.heap.Len()
}

func (q *PriorityQueue[T]) Clear() {
	q.heap.Clear()
}

func (q *PriorityQueue[T]) Enqueue(v T) {
	q.heap.Push(v)
}

func (q *PriorityQueue[T]) Dequeue() (T, bool) {
	return q.heap.Pop()
}

func (q *PriorityQueue[T]) Peek() (T, bool) {
	return q.heap.Peek()
}

func (q *PriorityQueue[T]) NextPriority() int64 {
	v, ok := q.Peek()
	if !ok {
		return -1
	}

	return v.Priority()
}
