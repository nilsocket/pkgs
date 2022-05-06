package pqueue

import (
	"testing"

	"github.com/nilsocket/pkgs/ds/heap/v2"
)

var input = []heap.Elem[int]{
	{Key: 5, Data: nil}, {Key: 3, Data: nil},
	{Key: 17, Data: nil}, {Key: 10, Data: nil},
	{Key: 84, Data: nil}, {Key: 19, Data: nil},
	{Key: 6, Data: nil}, {Key: 22, Data: nil},
	{Key: 9, Data: nil},
}

func TestRoot(t *testing.T) {
	inp := makeCopy(input)
	p := New(inp, heap.MaxHeap)

	assertRoot(t, p.Root().Key, 84)
}

func assertRoot(t *testing.T, gotRoot, wantRoot int) {
	t.Helper()
	if gotRoot != wantRoot {
		t.Errorf("gotRoot:%v, wantRoot:%v", gotRoot, wantRoot)
	}
}

func TestExtractRoot(t *testing.T) {
	inp := makeCopy(input)
	p := New(inp, heap.MaxHeap)

	assertRoot(t, p.ExtractRoot().Key, 84)
	assertRoot(t, p.ExtractRoot().Key, 22)
	assertRoot(t, p.ExtractRoot().Key, 19)
}

func TestModifyKey(t *testing.T) {
	inp := makeCopy(input)
	p := New(inp, heap.MaxHeap)
	p.ModifyKey(&inp[5], 1)
	heapValidate(t, p)
	// p.PrintTree()
}

func TestInsert(t *testing.T) {
	inp := makeCopy(input)
	p := New(inp, heap.MaxHeap)
	p.Insert(heap.Elem[int]{Key: 12})
	p.Insert(heap.Elem[int]{Key: 13})
	p.Insert(heap.Elem[int]{Key: 14})
	p.Insert(heap.Elem[int]{Key: 15})
	p.Insert(heap.Elem[int]{Key: 16})
	p.Insert(heap.Elem[int]{Key: 17})
	p.Insert(heap.Elem[int]{Key: 18})
	heapValidate(t, p)
	// p.PrintTree()
}

func makeCopy(x []heap.Elem[int]) []heap.Elem[int] {
	t := make([]heap.Elem[int], len(x))
	copy(t, x)
	return t
}

func heapValidate(t *testing.T, p PQueue[int]) {
	for i := heap.Parent(p.Size - 1); i > -1; i-- {
		l, r := heap.Left(i), heap.Right(i)

		var pn, ln, rn int
		pn = p.Elems[i].Key
		if l < p.Size {
			ln = p.Elems[l].Key
			if p.Type == heap.MaxHeap && ln > pn {
				t.Errorf("left node:%v is greater than parent node:%v", ln, pn)
			} else if p.Type == heap.MinHeap && ln < pn {
				t.Errorf("left node:%v is less than parent node:%v", ln, pn)
			}
		}
		if r < p.Size {
			rn = p.Elems[r].Key
			if p.Type == heap.MaxHeap && rn > pn {
				t.Errorf("right node:%v is greater than parent node:%v", rn, pn)
			} else if p.Type == heap.MinHeap && rn < pn {
				t.Errorf("right node:%v is less than parent node:%v", rn, pn)
			}
		}
	}
}
