package heap

import (
	"reflect"
	"testing"
)

var input = []Elem[int]{
	{5, nil}, {3, nil},
	{17, nil}, {10, nil},
	{84, nil}, {19, nil},
	{6, nil}, {22, nil},
	{9, nil},
}

func TestMaxHeap(t *testing.T) {
	want := []Elem[int]{
		{84, nil}, {22, nil},
		{19, nil}, {10, nil},
		{3, nil}, {17, nil},
		{6, nil}, {5, nil},
		{9, nil},
	}

	testMaxHeap(t, makeCopy(input), want)
}

func testMaxHeap(t *testing.T, input, want []Elem[int]) {
	t.Helper()

	inputCopy := makeCopy(input)

	New(input, MaxHeap)
	got := input

	assertEqual(t, inputCopy, got, want)
}

func assertEqual(t *testing.T, input, got, want []Elem[int]) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("input: %v\ngot : %v\nwant: %v", input, got, want)
	}
}

func makeCopy(x []Elem[int]) []Elem[int] {
	t := make([]Elem[int], len(x))
	copy(t, x)
	return t
}

func TestHeight(t *testing.T) {
	inp := makeCopy(input)
	h := New(inp, MaxHeap)
	got, want := h.Height(), 3
	if got != want {
		t.Errorf("height of the tree, got:%v, want:%v", got, want)
	}
}

func TestPrintTree(t *testing.T) {
	inp := makeCopy(input)
	h := New(inp, MaxHeap)
	h.PrintTree()
}
