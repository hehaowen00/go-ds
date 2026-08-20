package ds

import "errors"

type RingBuffer[T any] struct {
	data   []T
	count  int
	writer int
	reader int
}

var ErrCapacityInvalid = errors.New("capacity must be > 0")

func NewRingBuffer[T any](capacity int) (*RingBuffer[T], error) {
	if capacity <= 0 {
		return nil, ErrCapacityInvalid
	}

	return &RingBuffer[T]{
		data: make([]T, capacity),
	}, nil
}

func (r *RingBuffer[T]) Len() int {
	return r.count
}

// drops oldest if full
func (r *RingBuffer[T]) Push(v T) {
	r.data[r.writer] = v
	r.writer = (r.writer + 1) % len(r.data)

	if r.count == len(r.data) {
		r.reader = r.writer
		return
	} else {
		r.count++
	}
}

func (r *RingBuffer[T]) Pop() (T, bool) {
	if r.count == 0 {
		var empty T
		return empty, false
	}

	item := r.data[r.reader]
	var zero T
	r.data[r.reader] = zero
	r.reader = (r.reader + 1) % len(r.data)
	r.count--

	return item, true
}

func (r *RingBuffer[T]) Peek() (T, bool) {
	if r.count == 0 {
		var empty T
		return empty, false
	}

	item := r.data[r.reader]
	return item, true
}
