package ds

type Bloom[K any] struct {
}

func NewBloom[K any]() *Bloom[K] {
	return &Bloom[K]{}
}

func (b *Bloom[K]) Insert(item K) {
}

func (b *Bloom[K]) Exists(item K) {
}
