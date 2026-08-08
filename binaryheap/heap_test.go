package binaryheap

import (
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	bh := New(func(a, b int) bool {
		return a < b
	})

	for i := range 20 {
		bh.Push(i)
	}

	t.Log(bh.data)

	v, ok := bh.Pop()
	for ok {
		t.Log("out", v)
		v, ok = bh.Pop()
	}
}

const benchSize = 1024

func BenchmarkBinaryHeap(b *testing.B) {
	less := func(a, b int) bool {
		return a < b
	}

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		bh := New(less)
		for i := range benchSize {
			bh.Push(i)
		}
	}
}

func BenchmarkBinaryHeapCapacity(b *testing.B) {
	less := func(a, b int) bool {
		return a < b
	}

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		bh := NewCapacity(less, benchSize)
		for i := range benchSize {
			bh.Push(i)
		}
	}
}
