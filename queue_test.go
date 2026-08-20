package ds

import (
	"math/rand"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	for i := range 10 {
		q.Enqueue(i)
	}

	v, ok := q.Dequeue()
	for ok {
		t.Log(v)
		v, ok = q.Dequeue()
	}
}

type Item struct {
	id   int
	next int64
}

func (i Item) Priority() int64 {
	return i.next
}

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[Item]()

	for i := range 10 {
		n := rand.Intn(10)
		pq.Enqueue(Item{
			id:   i,
			next: time.Now().Add(time.Duration(n) * time.Second).Unix(),
		})
	}

	v, ok := pq.Dequeue()
	for ok {
		t.Log(v.id, v.next)
		v, ok = pq.Dequeue()
	}
}
