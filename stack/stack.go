package stack

type Stack[T any] struct {
	data []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func From[T any](data []T) *Stack[T] {
	return &Stack[T]{
		data: data,
	}
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Clear() {
	clear(s.data)
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var empty T
		return empty, false
	}

	last := len(s.data) - 1
	item := s.data[last]
	s.data = s.data[:last]

	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var empty T
		return empty, false
	}

	last := len(s.data) - 1
	item := s.data[last]

	return item, true
}
