package heap

import (
	"github.com/nilsocket/pkgs/ds/heap/v2"
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](elems []heap.Elem[T], ht heap.Type) {
	h := heap.New(elems, ht)

	for h.Size > 0 {
		h.Size--
		h.Elems[h.Size], h.Elems[0] = h.Elems[0], h.Elems[h.Size]
		h.HeapifyAt(0)
	}
}
