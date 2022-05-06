package heap

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return (i << 1) + 1
}

func Right(i int) int {
	return (i << 1) + 2
}

type Type uint8

const (
	MinHeap Type = iota
	MaxHeap
)

type Elem[T constraints.Ordered] struct {
	Key  T
	Data any
}

type Heap[T constraints.Ordered] struct {
	Elems []Elem[T]
	Size  int
	Type  Type
}

func New[T constraints.Ordered](elems []Elem[T], ht Type) *Heap[T] {
	h := Heap[T]{Elems: elems, Size: len(elems), Type: ht}
	return h.build()
}

func (h *Heap[T]) HeapifyAt(i int) {
	l, r := Left(i), Right(i)
	mi := i // max/min index

	if h.Type == MaxHeap {
		if l < h.Size && h.Elems[l].Key > h.Elems[i].Key {
			mi = l
		}
		if r < h.Size && h.Elems[r].Key > h.Elems[mi].Key {
			mi = r
		}
	} else {
		if l < h.Size && h.Elems[l].Key < h.Elems[i].Key {
			mi = l
		}
		if l < h.Size && h.Elems[r].Key < h.Elems[mi].Key {
			mi = r
		}
	}

	if mi != i {
		h.Elems[i], h.Elems[mi] = h.Elems[mi], h.Elems[i]
		h.HeapifyAt(mi)
	}
}

func (h *Heap[T]) build() *Heap[T] {
	// heapify from
	hf := Parent(h.Size - 1) // last node which have a child

	for i := hf; i > -1; i-- {
		h.HeapifyAt(i)
	}
	return h
}

func (h *Heap[T]) Rebuild() {
	h.build()
}

func (h *Heap[T]) Height() int {
	return int(math.Log2(float64(h.Size)))
}
