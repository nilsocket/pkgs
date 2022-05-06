package heap

import (
	"github.com/nilsocket/pkgs/ds/heap/v0"
	"golang.org/x/exp/constraints"
)

func MaxSort[T constraints.Ordered](a []T) {
	h := heap.MaxHeap(a)

	for i := h.Size - 1; i > 0; i-- {
		h.Data[i], h.Data[0] = h.Data[0], h.Data[i]
		h.Size--
		h.MaxHeapifyAt(0)
	}
}

func MinSort[T constraints.Ordered](a []T) {
	h := heap.MinHeap(a)

	for i := h.Size - 1; i > 0; i-- {
		h.Data[i], h.Data[0] = h.Data[0], h.Data[i]
		h.Size--
		h.MinHeapifyAt(0)
	}
}
