package ring

import "testing"

func TestRing(t *testing.T) {
	ringBuf, err := New[int](16)
	if err != nil {
		t.FailNow()
	}
	for i := range 64 {
		ringBuf.Push(i)
	}

	t.Log(ringBuf.data)

	v, ok := ringBuf.Pop()
	for ok {
		t.Log(v)
		v, ok = ringBuf.Pop()
	}
}
