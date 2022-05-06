package heap

import (
	"reflect"
	"testing"
)

type MyType []struct {
	key       int
	someData  string
	someData2 any
}

func (m MyType) Len() int      { return len(m) }
func (m MyType) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

// >, maxHeap
// <, minHeap
func (m MyType) Less(i, j int) bool { return m[i].key > m[j].key }

var input = MyType{
	{5, "x", "y"},
	{3, "x", "y"},
	{17, "x", "y"},
	{10, "x", "y"},
	{84, "x", "y"},
	{19, "x", "y"},
	{6, "x", "y"},
	{22, "x", "y"},
	{9, "x", "y"},
}

var want = MyType{
	{84, "x", "y"},
	{22, "x", "y"},
	{19, "x", "y"},
	{10, "x", "y"},
	{3, "x", "y"},
	{17, "x", "y"},
	{6, "x", "y"},
	{5, "x", "y"},
	{9, "x", "y"},
}

func TestMaxHeap(t *testing.T) {
	inputCopy := makeCopy(input)
	New(input)
	got := input
	assertEqual(t, inputCopy, got, want)
}

func assertEqual(t *testing.T, input, got, want MyType) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("input: %v\ngot : %v\nwant: %v", input, got, want)
	}
}

func makeCopy(x MyType) MyType {
	t := make(MyType, x.Len())
	copy(t, x)
	return t
}
