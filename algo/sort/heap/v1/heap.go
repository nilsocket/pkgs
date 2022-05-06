package heap

import (
	"sort"

	"github.com/nilsocket/pkgs/ds/heap/v1"
)

func Sort(data sort.Interface) {
	h := heap.New(data)
	for ; h.Size > 1; h.Size-- {
		h.HeapifyAt(0)
		h.Data.Swap(0, h.Size-1)
	}
	h.Size = data.Len()
}
