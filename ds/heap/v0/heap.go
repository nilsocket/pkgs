package heap

import "golang.org/x/exp/constraints"

func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return (2 * i) + 1
}

func Right(i int) int {
	return (2 * i) + 2
}

type Heap[T constraints.Ordered] struct {
	Data []T
	Size int
}

// max heap

func MaxHeap[T constraints.Ordered](a []T) *Heap[T] {
	h := Heap[T]{a, len(a)}
	for i := Parent(h.Size - 1); i >= 0; i-- {
		h.maxHeapifyAt(i)
	}
	return &h
}

func (h *Heap[T]) MaxHeapifyAt(i int) {
	h.maxHeapifyAt(i)
}

func (h *Heap[T]) maxHeapifyAt(i int) {
	l, r := Left(i), Right(i)
	li := i // largest index

	if l < h.Size && h.Data[l] > h.Data[i] {
		li = l
	}
	if r < h.Size && h.Data[r] > h.Data[li] {
		li = r
	}

	if li != i {
		h.Data[li], h.Data[i] = h.Data[i], h.Data[li]
		h.maxHeapifyAt(li)
	}
}

// min heap

func MinHeap[T constraints.Ordered](a []T) *Heap[T] {
	h := Heap[T]{a, len(a)}
	for i := Parent(h.Size - 1); i >= 0; i-- {
		h.minHeapifyAt(i)
	}
	return &h
}

func (h *Heap[T]) MinHeapifyAt(i int) {
	h.minHeapifyAt(i)
}

func (h *Heap[T]) minHeapifyAt(i int) {
	l, r := Left(i), Right(i)
	si := i // smallest index

	if l < h.Size && h.Data[l] < h.Data[i] {
		si = l
	}
	if r < h.Size && h.Data[r] < h.Data[si] {
		si = r
	}

	if si != i {
		h.Data[si], h.Data[i] = h.Data[i], h.Data[si]
		h.minHeapifyAt(si)
	}
}
