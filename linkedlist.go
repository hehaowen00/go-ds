package ds

import (
	"iter"
)

type SinglyLinkedList[T any] struct {
	head *linkedlistnode[T]
	tail *linkedlistnode[T]
	size int
}

type linkedlistnode[T any] struct {
	value T
	next  *linkedlistnode[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) Len() int {
	return l.size
}

func (l *SinglyLinkedList[T]) Index(idx int) (T, bool) {
	if idx < 0 || l.head == nil {
		var empty T
		return empty, false
	}

	index := 0
	cursor := l.head

	for {
		if index == idx {
			return cursor.value, true
		}

		if cursor.next == nil {
			var empty T
			return empty, false
		}

		cursor = cursor.next
		index += 1
	}
}

func (l *SinglyLinkedList[T]) Remove(idx int) (T, bool) {
	if idx < 0 || l.head == nil {
		var empty T
		return empty, false
	}

	index := 0
	cursor := l.head
	var last *linkedlistnode[T]

	for {
		if index == idx {
			if last == nil {
				l.head = cursor.next
			} else {
				last.next = cursor.next
			}

			if cursor == l.tail {
				l.tail = last
			}

			l.size -= 1
			return cursor.value, true
		}

		if cursor.next == nil {
			var empty T
			return empty, false
		}

		last = cursor
		cursor = cursor.next
		index += 1
	}
}

func (l *SinglyLinkedList[T]) PushFront(val T) {
	node := linkedlistnode[T]{
		value: val,
		next:  l.head,
	}

	l.head = &node
	if l.tail == nil {
		l.tail = l.head
	}
	l.size += 1
}

func (l *SinglyLinkedList[T]) PushBack(val T) {
	n := &linkedlistnode[T]{
		value: val,
	}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}

	l.tail = n
	l.size += 1
}

func (l *SinglyLinkedList[T]) PopBack() (T, bool) {
	if l.size == 0 {
		var empty T
		return empty, false
	}

	cursor := l.head
	var last *linkedlistnode[T]

	for cursor.next != nil {
		last = cursor
		cursor = cursor.next
	}

	l.size -= 1

	if last != nil {
		last.next = nil
		l.tail = last

		return cursor.value, true
	}

	l.head = nil
	l.tail = nil

	return cursor.value, true
}

func (l *SinglyLinkedList[T]) PopFront() (T, bool) {
	if l.head == nil {
		var empty T
		return empty, false
	}

	result := l.head.value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	}

	l.size -= 1

	return result, true
}

func (l *SinglyLinkedList[T]) Iterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		cursor := l.head
		i := 0

		for cursor != nil {
			if !yield(i, cursor.value) {
				return
			}

			cursor = cursor.next
			i += 1
		}
	}
}

type DoublyLinkedList[T any] struct {
	head *doublylinkedlistnode[T]
	tail *doublylinkedlistnode[T]
	size int
}

type doublylinkedlistnode[T any] struct {
	value T
	prev  *doublylinkedlistnode[T]
	next  *doublylinkedlistnode[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.size
}

func (l *DoublyLinkedList[T]) Index(idx int) (T, bool) {
	if idx < 0 || l.head == nil {
		var empty T
		return empty, false
	}

	index := 0
	cursor := l.head

	for {
		if index == idx {
			return cursor.value, true
		}

		if cursor.next == nil {
			var empty T
			return empty, false
		}

		cursor = cursor.next
		index += 1
	}
}

func (l *DoublyLinkedList[T]) Remove(idx int) (T, bool) {
	if idx < 0 || l.head == nil {
		var empty T
		return empty, false
	}

	index := 0
	cursor := l.head
	var last *doublylinkedlistnode[T]

	for {
		if index == idx {
			if last == nil {
				l.head = cursor.next
			} else {
				last.next = cursor.next
			}

			if cursor == l.tail {
				l.tail = last
			}

			l.size -= 1
			return cursor.value, true
		}

		if cursor.next == nil {
			var empty T
			return empty, false
		}

		last = cursor
		cursor = cursor.next
		index += 1
	}
}

func (l *DoublyLinkedList[T]) PushFront(val T) {
	node := doublylinkedlistnode[T]{
		value: val,
		next:  l.head,
	}

	l.head = &node

	if l.tail == nil {
		l.tail = l.head
	}

	l.size += 1
}

func (l *DoublyLinkedList[T]) PushBack(val T) {
	n := &doublylinkedlistnode[T]{
		value: val,
	}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}

	l.tail = n
	l.size += 1
}

func (l *DoublyLinkedList[T]) PopBack() (T, bool) {
	if l.size == 0 {
		var empty T
		return empty, false
	}

	cursor := l.head
	var last *doublylinkedlistnode[T]

	for cursor.next != nil {
		last = cursor
		cursor = cursor.next
	}

	l.size -= 1

	if last != nil {
		last.next = nil
		l.tail = last

		return cursor.value, true
	}

	l.head = nil
	l.tail = nil

	return cursor.value, true
}

func (l *DoublyLinkedList[T]) PopFront() (T, bool) {
	if l.head == nil {
		var empty T
		return empty, false
	}

	result := l.head.value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	}

	l.size -= 1

	return result, true
}

func (l *DoublyLinkedList[T]) Iterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		cursor := l.head
		i := 0

		for cursor != nil {
			if !yield(i, cursor.value) {
				return
			}

			cursor = cursor.next
			i += 1
		}
	}
}
