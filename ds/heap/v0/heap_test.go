package heap

import (
	"reflect"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	input := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	want := []int{84, 22, 19, 10, 3, 17, 6, 5, 9}

	testMaxHeap(t, input, want)
}

func testMaxHeap(t *testing.T, input, want []int) {
	t.Helper()

	inputCopy := makeCopy(input)

	MaxHeap(input)
	got := input

	assertEqual(t, inputCopy, got, want)
}

func assertEqual(t *testing.T, input, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("input: %v\ngot : %v\nwant: %v", input, got, want)
	}
}

func makeCopy(x []int) []int {
	t := make([]int, len(x))
	copy(t, x)
	return t
}
