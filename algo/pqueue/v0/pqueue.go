package pqueue

import (
	"unsafe"

	"github.com/nilsocket/pkgs/ds/heap/v2"
	"golang.org/x/exp/constraints"
)

type PQueue[T constraints.Ordered] struct {
	*heap.Heap[T]
}

func New[T constraints.Ordered](elems []heap.Elem[T], ht heap.Type) PQueue[T] {
	return PQueue[T]{heap.New(elems, ht)}
}

func (p PQueue[T]) Root() *heap.Elem[T] {
	if p.Size > 0 {
		return &p.Elems[0]
	}
	return nil
}

func (p PQueue[T]) ExtractRoot() *heap.Elem[T] {
	if p.Size > 0 {
		p.Size--
		p.Elems[p.Size], p.Elems[0] = p.Elems[0], p.Elems[p.Size]
		p.HeapifyAt(0)
		return &p.Elems[p.Size]
	}
	return nil
}

func (p PQueue[T]) Insert(e heap.Elem[T]) {
	p.Size++
	p.Elems = append(p.Elems, e)
	for i := heap.Parent(p.Size - 1); i > 0; i = heap.Parent(i) {
		p.HeapifyAt(i)
	}
}

// ModifyKey of given elem e with nk
// caveat: for now e should belong to the same array
// as provided to New()
func (p PQueue[T]) ModifyKey(e *heap.Elem[T], nk T) {
	i := int((uintptr(unsafe.Pointer(e)) - uintptr(unsafe.Pointer(&p.Elems[0]))) / unsafe.Sizeof(p.Elems[0]))
	p.Elems[i].Key = nk

	for ep := heap.Parent(i); ep > 0; ep = heap.Parent(ep) {
		p.HeapifyAt(ep)
	}
}
