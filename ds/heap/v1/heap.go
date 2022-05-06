package heap

import (
	"sort"
)

func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return (2 * i) + 1
}

func Right(i int) int {
	return (2 * i) + 2
}

type Heap struct {
	Data sort.Interface
	Size int
}

func New(data sort.Interface) *Heap {
	return heap(data)
}

func heap(data sort.Interface) *Heap {
	h := Heap{data, data.Len()}
	for i := Parent(h.Size - 1); i >= 0; i-- {
		h.HeapifyAt(i)
	}
	return &h
}

func (h *Heap) HeapifyAt(i int) {
	l, r := Left(i), Right(i)
	li := i // largest index

	if l < h.Size && h.Data.Less(l, i) {
		li = l
	}
	if r < h.Size && h.Data.Less(r, li) {
		li = r
	}

	if li != i {
		h.Data.Swap(li, i)
		h.HeapifyAt(li)
	}
}
